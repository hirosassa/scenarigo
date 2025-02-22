package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Yamashou/gqlgenc/clientv2"

	"github.com/zoncoen/scenarigo/scripts/cross-build/gen"
)

var (
	token            = os.Getenv("GITHUB_TOKEN")
	releaseVer       = os.Getenv("RELEASE_VERSION")
	ver              = os.Getenv("GO_VERSION")
	rootDir          = os.Getenv("PJ_ROOT")
	errImageNotFound = errors.New("image not found")
)

func main() {
	if err := release(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func release() error {
	tag, err := imageTag(ver, token)
	if err != nil {
		if errors.Is(err, errImageNotFound) {
			fmt.Printf("golang-cross:v%s: %s\n", ver, err)
			return nil
		}
		return fmt.Errorf("failed to get image tag: %w", err)
	}
	if err := build(ver, tag); err != nil {
		return fmt.Errorf("failed to build: %w", err)
	}
	return nil
}

func imageTag(ver, token string) (string, error) {
	// HACK: golang-cross:v1.17.5-4 does not work
	if ver == "1.17.5" {
		return "v1.17.5-0", nil
	}
	github := &gen.Client{
		Client: clientv2.NewClient(http.DefaultClient, "https://api.github.com/graphql", func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return next(ctx, req, gqlInfo, res)
		}),
	}

	ctx := context.Background()
	first := 100
	getTags, err := github.GetTags(ctx, "gythialy", "golang-cross", &first)
	if err != nil {
		if handledError, ok := err.(*clientv2.ErrorResponse); ok {
			return "", fmt.Errorf("handled error: %sn", handledError.Error())
		}
		return "", fmt.Errorf("unhandled error: %s", err.Error())
	}
	v := fmt.Sprintf("v%s", ver)
	prefix := fmt.Sprintf("%s-", v)
	for _, node := range getTags.Repository.Refs.Nodes {
		if node.Name == v {
			return node.Name, nil
		}
		if strings.HasPrefix(node.Name, prefix) {
			return node.Name, nil
		}
	}

	return "", errImageNotFound
}

//go:embed templates/goreleaser.yml.tmpl
var tmplBytes []byte

func build(ver, tag string) error {
	if err := os.Mkdir(fmt.Sprintf("%s/assets", rootDir), 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpl := template.New("config").Delims("<<", ">>")
	tmpl, err := tmpl.Parse(string(tmplBytes))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/.goreleaser.yml", rootDir), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		return fmt.Errorf("failed to open .goreleaser.yml: %w", err)
	}
	defer f.Close()
	if err := tmpl.Execute(f, map[string]interface{}{
		"GoVersion": ver,
	}); err != nil {
		return fmt.Errorf("failed to create .goreleaser.yml: %w", err)
	}
	if err := goreleaser(ver, tag); err != nil {
		return fmt.Errorf("goreleaser failed (go%s, golang-cross:%s): %w", ver, tag, err)
	}

	return nil
}

func goreleaser(ver, tag string) error {
	out, err := exec.Command(
		"docker", "run", "--rm", "--privileged",
		"-v", fmt.Sprintf("%s:/scenarigo", rootDir),
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-w", "/scenarigo",
		fmt.Sprintf("ghcr.io/gythialy/golang-cross:%s", tag),
		"--skip-publish", "--rm-dist",
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s:\n%s", err, out)
	}
	return nil
}
