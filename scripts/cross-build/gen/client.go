// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, interceptors ...clientv2.RequestInterceptor) *Client {
	return &Client{Client: clientv2.NewClient(cli, baseURL, interceptors...)}
}

type Query struct {
	CodeOfConduct                            *CodeOfConduct                     "json:\"codeOfConduct\" graphql:\"codeOfConduct\""
	CodesOfConduct                           []*CodeOfConduct                   "json:\"codesOfConduct\" graphql:\"codesOfConduct\""
	Enterprise                               *Enterprise                        "json:\"enterprise\" graphql:\"enterprise\""
	EnterpriseAdministratorInvitation        *EnterpriseAdministratorInvitation "json:\"enterpriseAdministratorInvitation\" graphql:\"enterpriseAdministratorInvitation\""
	EnterpriseAdministratorInvitationByToken *EnterpriseAdministratorInvitation "json:\"enterpriseAdministratorInvitationByToken\" graphql:\"enterpriseAdministratorInvitationByToken\""
	License                                  *License                           "json:\"license\" graphql:\"license\""
	Licenses                                 []*License                         "json:\"licenses\" graphql:\"licenses\""
	MarketplaceCategories                    []*MarketplaceCategory             "json:\"marketplaceCategories\" graphql:\"marketplaceCategories\""
	MarketplaceCategory                      *MarketplaceCategory               "json:\"marketplaceCategory\" graphql:\"marketplaceCategory\""
	MarketplaceListing                       *MarketplaceListing                "json:\"marketplaceListing\" graphql:\"marketplaceListing\""
	MarketplaceListings                      MarketplaceListingConnection       "json:\"marketplaceListings\" graphql:\"marketplaceListings\""
	Meta                                     GitHubMetadata                     "json:\"meta\" graphql:\"meta\""
	Node                                     Node                               "json:\"node\" graphql:\"node\""
	Nodes                                    []Node                             "json:\"nodes\" graphql:\"nodes\""
	Organization                             *Organization                      "json:\"organization\" graphql:\"organization\""
	RateLimit                                *RateLimit                         "json:\"rateLimit\" graphql:\"rateLimit\""
	Relay                                    *Query                             "json:\"relay\" graphql:\"relay\""
	Repository                               *Repository                        "json:\"repository\" graphql:\"repository\""
	RepositoryOwner                          RepositoryOwner                    "json:\"repositoryOwner\" graphql:\"repositoryOwner\""
	Resource                                 UniformResourceLocatable           "json:\"resource\" graphql:\"resource\""
	Search                                   SearchResultItemConnection         "json:\"search\" graphql:\"search\""
	SecurityAdvisories                       SecurityAdvisoryConnection         "json:\"securityAdvisories\" graphql:\"securityAdvisories\""
	SecurityAdvisory                         *SecurityAdvisory                  "json:\"securityAdvisory\" graphql:\"securityAdvisory\""
	SecurityVulnerabilities                  SecurityVulnerabilityConnection    "json:\"securityVulnerabilities\" graphql:\"securityVulnerabilities\""
	Sponsorables                             SponsorableItemConnection          "json:\"sponsorables\" graphql:\"sponsorables\""
	SponsorsListing                          *SponsorsListing                   "json:\"sponsorsListing\" graphql:\"sponsorsListing\""
	Topic                                    *Topic                             "json:\"topic\" graphql:\"topic\""
	User                                     *User                              "json:\"user\" graphql:\"user\""
	Viewer                                   User                               "json:\"viewer\" graphql:\"viewer\""
}
type Mutation struct {
	AcceptEnterpriseAdministratorInvitation                     *AcceptEnterpriseAdministratorInvitationPayload                     "json:\"acceptEnterpriseAdministratorInvitation\" graphql:\"acceptEnterpriseAdministratorInvitation\""
	AcceptTopicSuggestion                                       *AcceptTopicSuggestionPayload                                       "json:\"acceptTopicSuggestion\" graphql:\"acceptTopicSuggestion\""
	AddAssigneesToAssignable                                    *AddAssigneesToAssignablePayload                                    "json:\"addAssigneesToAssignable\" graphql:\"addAssigneesToAssignable\""
	AddComment                                                  *AddCommentPayload                                                  "json:\"addComment\" graphql:\"addComment\""
	AddDiscussionComment                                        *AddDiscussionCommentPayload                                        "json:\"addDiscussionComment\" graphql:\"addDiscussionComment\""
	AddEnterpriseSupportEntitlement                             *AddEnterpriseSupportEntitlementPayload                             "json:\"addEnterpriseSupportEntitlement\" graphql:\"addEnterpriseSupportEntitlement\""
	AddLabelsToLabelable                                        *AddLabelsToLabelablePayload                                        "json:\"addLabelsToLabelable\" graphql:\"addLabelsToLabelable\""
	AddProjectCard                                              *AddProjectCardPayload                                              "json:\"addProjectCard\" graphql:\"addProjectCard\""
	AddProjectColumn                                            *AddProjectColumnPayload                                            "json:\"addProjectColumn\" graphql:\"addProjectColumn\""
	AddPullRequestReview                                        *AddPullRequestReviewPayload                                        "json:\"addPullRequestReview\" graphql:\"addPullRequestReview\""
	AddPullRequestReviewComment                                 *AddPullRequestReviewCommentPayload                                 "json:\"addPullRequestReviewComment\" graphql:\"addPullRequestReviewComment\""
	AddPullRequestReviewThread                                  *AddPullRequestReviewThreadPayload                                  "json:\"addPullRequestReviewThread\" graphql:\"addPullRequestReviewThread\""
	AddReaction                                                 *AddReactionPayload                                                 "json:\"addReaction\" graphql:\"addReaction\""
	AddStar                                                     *AddStarPayload                                                     "json:\"addStar\" graphql:\"addStar\""
	AddUpvote                                                   *AddUpvotePayload                                                   "json:\"addUpvote\" graphql:\"addUpvote\""
	AddVerifiableDomain                                         *AddVerifiableDomainPayload                                         "json:\"addVerifiableDomain\" graphql:\"addVerifiableDomain\""
	ApproveDeployments                                          *ApproveDeploymentsPayload                                          "json:\"approveDeployments\" graphql:\"approveDeployments\""
	ApproveVerifiableDomain                                     *ApproveVerifiableDomainPayload                                     "json:\"approveVerifiableDomain\" graphql:\"approveVerifiableDomain\""
	ArchiveRepository                                           *ArchiveRepositoryPayload                                           "json:\"archiveRepository\" graphql:\"archiveRepository\""
	CancelEnterpriseAdminInvitation                             *CancelEnterpriseAdminInvitationPayload                             "json:\"cancelEnterpriseAdminInvitation\" graphql:\"cancelEnterpriseAdminInvitation\""
	ChangeUserStatus                                            *ChangeUserStatusPayload                                            "json:\"changeUserStatus\" graphql:\"changeUserStatus\""
	ClearLabelsFromLabelable                                    *ClearLabelsFromLabelablePayload                                    "json:\"clearLabelsFromLabelable\" graphql:\"clearLabelsFromLabelable\""
	CloneProject                                                *CloneProjectPayload                                                "json:\"cloneProject\" graphql:\"cloneProject\""
	CloneTemplateRepository                                     *CloneTemplateRepositoryPayload                                     "json:\"cloneTemplateRepository\" graphql:\"cloneTemplateRepository\""
	CloseIssue                                                  *CloseIssuePayload                                                  "json:\"closeIssue\" graphql:\"closeIssue\""
	ClosePullRequest                                            *ClosePullRequestPayload                                            "json:\"closePullRequest\" graphql:\"closePullRequest\""
	ConvertProjectCardNoteToIssue                               *ConvertProjectCardNoteToIssuePayload                               "json:\"convertProjectCardNoteToIssue\" graphql:\"convertProjectCardNoteToIssue\""
	ConvertPullRequestToDraft                                   *ConvertPullRequestToDraftPayload                                   "json:\"convertPullRequestToDraft\" graphql:\"convertPullRequestToDraft\""
	CreateBranchProtectionRule                                  *CreateBranchProtectionRulePayload                                  "json:\"createBranchProtectionRule\" graphql:\"createBranchProtectionRule\""
	CreateCheckRun                                              *CreateCheckRunPayload                                              "json:\"createCheckRun\" graphql:\"createCheckRun\""
	CreateCheckSuite                                            *CreateCheckSuitePayload                                            "json:\"createCheckSuite\" graphql:\"createCheckSuite\""
	CreateDiscussion                                            *CreateDiscussionPayload                                            "json:\"createDiscussion\" graphql:\"createDiscussion\""
	CreateEnterpriseOrganization                                *CreateEnterpriseOrganizationPayload                                "json:\"createEnterpriseOrganization\" graphql:\"createEnterpriseOrganization\""
	CreateEnvironment                                           *CreateEnvironmentPayload                                           "json:\"createEnvironment\" graphql:\"createEnvironment\""
	CreateIPAllowListEntry                                      *CreateIPAllowListEntryPayload                                      "json:\"createIpAllowListEntry\" graphql:\"createIpAllowListEntry\""
	CreateIssue                                                 *CreateIssuePayload                                                 "json:\"createIssue\" graphql:\"createIssue\""
	CreateProject                                               *CreateProjectPayload                                               "json:\"createProject\" graphql:\"createProject\""
	CreatePullRequest                                           *CreatePullRequestPayload                                           "json:\"createPullRequest\" graphql:\"createPullRequest\""
	CreateRef                                                   *CreateRefPayload                                                   "json:\"createRef\" graphql:\"createRef\""
	CreateRepository                                            *CreateRepositoryPayload                                            "json:\"createRepository\" graphql:\"createRepository\""
	CreateSponsorship                                           *CreateSponsorshipPayload                                           "json:\"createSponsorship\" graphql:\"createSponsorship\""
	CreateTeamDiscussion                                        *CreateTeamDiscussionPayload                                        "json:\"createTeamDiscussion\" graphql:\"createTeamDiscussion\""
	CreateTeamDiscussionComment                                 *CreateTeamDiscussionCommentPayload                                 "json:\"createTeamDiscussionComment\" graphql:\"createTeamDiscussionComment\""
	DeclineTopicSuggestion                                      *DeclineTopicSuggestionPayload                                      "json:\"declineTopicSuggestion\" graphql:\"declineTopicSuggestion\""
	DeleteBranchProtectionRule                                  *DeleteBranchProtectionRulePayload                                  "json:\"deleteBranchProtectionRule\" graphql:\"deleteBranchProtectionRule\""
	DeleteDeployment                                            *DeleteDeploymentPayload                                            "json:\"deleteDeployment\" graphql:\"deleteDeployment\""
	DeleteDiscussion                                            *DeleteDiscussionPayload                                            "json:\"deleteDiscussion\" graphql:\"deleteDiscussion\""
	DeleteDiscussionComment                                     *DeleteDiscussionCommentPayload                                     "json:\"deleteDiscussionComment\" graphql:\"deleteDiscussionComment\""
	DeleteEnvironment                                           *DeleteEnvironmentPayload                                           "json:\"deleteEnvironment\" graphql:\"deleteEnvironment\""
	DeleteIPAllowListEntry                                      *DeleteIPAllowListEntryPayload                                      "json:\"deleteIpAllowListEntry\" graphql:\"deleteIpAllowListEntry\""
	DeleteIssue                                                 *DeleteIssuePayload                                                 "json:\"deleteIssue\" graphql:\"deleteIssue\""
	DeleteIssueComment                                          *DeleteIssueCommentPayload                                          "json:\"deleteIssueComment\" graphql:\"deleteIssueComment\""
	DeleteProject                                               *DeleteProjectPayload                                               "json:\"deleteProject\" graphql:\"deleteProject\""
	DeleteProjectCard                                           *DeleteProjectCardPayload                                           "json:\"deleteProjectCard\" graphql:\"deleteProjectCard\""
	DeleteProjectColumn                                         *DeleteProjectColumnPayload                                         "json:\"deleteProjectColumn\" graphql:\"deleteProjectColumn\""
	DeletePullRequestReview                                     *DeletePullRequestReviewPayload                                     "json:\"deletePullRequestReview\" graphql:\"deletePullRequestReview\""
	DeletePullRequestReviewComment                              *DeletePullRequestReviewCommentPayload                              "json:\"deletePullRequestReviewComment\" graphql:\"deletePullRequestReviewComment\""
	DeleteRef                                                   *DeleteRefPayload                                                   "json:\"deleteRef\" graphql:\"deleteRef\""
	DeleteTeamDiscussion                                        *DeleteTeamDiscussionPayload                                        "json:\"deleteTeamDiscussion\" graphql:\"deleteTeamDiscussion\""
	DeleteTeamDiscussionComment                                 *DeleteTeamDiscussionCommentPayload                                 "json:\"deleteTeamDiscussionComment\" graphql:\"deleteTeamDiscussionComment\""
	DeleteVerifiableDomain                                      *DeleteVerifiableDomainPayload                                      "json:\"deleteVerifiableDomain\" graphql:\"deleteVerifiableDomain\""
	DisablePullRequestAutoMerge                                 *DisablePullRequestAutoMergePayload                                 "json:\"disablePullRequestAutoMerge\" graphql:\"disablePullRequestAutoMerge\""
	DismissPullRequestReview                                    *DismissPullRequestReviewPayload                                    "json:\"dismissPullRequestReview\" graphql:\"dismissPullRequestReview\""
	EnablePullRequestAutoMerge                                  *EnablePullRequestAutoMergePayload                                  "json:\"enablePullRequestAutoMerge\" graphql:\"enablePullRequestAutoMerge\""
	FollowUser                                                  *FollowUserPayload                                                  "json:\"followUser\" graphql:\"followUser\""
	InviteEnterpriseAdmin                                       *InviteEnterpriseAdminPayload                                       "json:\"inviteEnterpriseAdmin\" graphql:\"inviteEnterpriseAdmin\""
	LinkRepositoryToProject                                     *LinkRepositoryToProjectPayload                                     "json:\"linkRepositoryToProject\" graphql:\"linkRepositoryToProject\""
	LockLockable                                                *LockLockablePayload                                                "json:\"lockLockable\" graphql:\"lockLockable\""
	MarkDiscussionCommentAsAnswer                               *MarkDiscussionCommentAsAnswerPayload                               "json:\"markDiscussionCommentAsAnswer\" graphql:\"markDiscussionCommentAsAnswer\""
	MarkFileAsViewed                                            *MarkFileAsViewedPayload                                            "json:\"markFileAsViewed\" graphql:\"markFileAsViewed\""
	MarkPullRequestReadyForReview                               *MarkPullRequestReadyForReviewPayload                               "json:\"markPullRequestReadyForReview\" graphql:\"markPullRequestReadyForReview\""
	MergeBranch                                                 *MergeBranchPayload                                                 "json:\"mergeBranch\" graphql:\"mergeBranch\""
	MergePullRequest                                            *MergePullRequestPayload                                            "json:\"mergePullRequest\" graphql:\"mergePullRequest\""
	MinimizeComment                                             *MinimizeCommentPayload                                             "json:\"minimizeComment\" graphql:\"minimizeComment\""
	MoveProjectCard                                             *MoveProjectCardPayload                                             "json:\"moveProjectCard\" graphql:\"moveProjectCard\""
	MoveProjectColumn                                           *MoveProjectColumnPayload                                           "json:\"moveProjectColumn\" graphql:\"moveProjectColumn\""
	PinIssue                                                    *PinIssuePayload                                                    "json:\"pinIssue\" graphql:\"pinIssue\""
	RegenerateEnterpriseIdentityProviderRecoveryCodes           *RegenerateEnterpriseIdentityProviderRecoveryCodesPayload           "json:\"regenerateEnterpriseIdentityProviderRecoveryCodes\" graphql:\"regenerateEnterpriseIdentityProviderRecoveryCodes\""
	RegenerateVerifiableDomainToken                             *RegenerateVerifiableDomainTokenPayload                             "json:\"regenerateVerifiableDomainToken\" graphql:\"regenerateVerifiableDomainToken\""
	RejectDeployments                                           *RejectDeploymentsPayload                                           "json:\"rejectDeployments\" graphql:\"rejectDeployments\""
	RemoveAssigneesFromAssignable                               *RemoveAssigneesFromAssignablePayload                               "json:\"removeAssigneesFromAssignable\" graphql:\"removeAssigneesFromAssignable\""
	RemoveEnterpriseAdmin                                       *RemoveEnterpriseAdminPayload                                       "json:\"removeEnterpriseAdmin\" graphql:\"removeEnterpriseAdmin\""
	RemoveEnterpriseIdentityProvider                            *RemoveEnterpriseIdentityProviderPayload                            "json:\"removeEnterpriseIdentityProvider\" graphql:\"removeEnterpriseIdentityProvider\""
	RemoveEnterpriseOrganization                                *RemoveEnterpriseOrganizationPayload                                "json:\"removeEnterpriseOrganization\" graphql:\"removeEnterpriseOrganization\""
	RemoveEnterpriseSupportEntitlement                          *RemoveEnterpriseSupportEntitlementPayload                          "json:\"removeEnterpriseSupportEntitlement\" graphql:\"removeEnterpriseSupportEntitlement\""
	RemoveLabelsFromLabelable                                   *RemoveLabelsFromLabelablePayload                                   "json:\"removeLabelsFromLabelable\" graphql:\"removeLabelsFromLabelable\""
	RemoveOutsideCollaborator                                   *RemoveOutsideCollaboratorPayload                                   "json:\"removeOutsideCollaborator\" graphql:\"removeOutsideCollaborator\""
	RemoveReaction                                              *RemoveReactionPayload                                              "json:\"removeReaction\" graphql:\"removeReaction\""
	RemoveStar                                                  *RemoveStarPayload                                                  "json:\"removeStar\" graphql:\"removeStar\""
	RemoveUpvote                                                *RemoveUpvotePayload                                                "json:\"removeUpvote\" graphql:\"removeUpvote\""
	ReopenIssue                                                 *ReopenIssuePayload                                                 "json:\"reopenIssue\" graphql:\"reopenIssue\""
	ReopenPullRequest                                           *ReopenPullRequestPayload                                           "json:\"reopenPullRequest\" graphql:\"reopenPullRequest\""
	RequestReviews                                              *RequestReviewsPayload                                              "json:\"requestReviews\" graphql:\"requestReviews\""
	RerequestCheckSuite                                         *RerequestCheckSuitePayload                                         "json:\"rerequestCheckSuite\" graphql:\"rerequestCheckSuite\""
	ResolveReviewThread                                         *ResolveReviewThreadPayload                                         "json:\"resolveReviewThread\" graphql:\"resolveReviewThread\""
	SetEnterpriseIdentityProvider                               *SetEnterpriseIdentityProviderPayload                               "json:\"setEnterpriseIdentityProvider\" graphql:\"setEnterpriseIdentityProvider\""
	SetOrganizationInteractionLimit                             *SetOrganizationInteractionLimitPayload                             "json:\"setOrganizationInteractionLimit\" graphql:\"setOrganizationInteractionLimit\""
	SetRepositoryInteractionLimit                               *SetRepositoryInteractionLimitPayload                               "json:\"setRepositoryInteractionLimit\" graphql:\"setRepositoryInteractionLimit\""
	SetUserInteractionLimit                                     *SetUserInteractionLimitPayload                                     "json:\"setUserInteractionLimit\" graphql:\"setUserInteractionLimit\""
	SubmitPullRequestReview                                     *SubmitPullRequestReviewPayload                                     "json:\"submitPullRequestReview\" graphql:\"submitPullRequestReview\""
	TransferIssue                                               *TransferIssuePayload                                               "json:\"transferIssue\" graphql:\"transferIssue\""
	UnarchiveRepository                                         *UnarchiveRepositoryPayload                                         "json:\"unarchiveRepository\" graphql:\"unarchiveRepository\""
	UnfollowUser                                                *UnfollowUserPayload                                                "json:\"unfollowUser\" graphql:\"unfollowUser\""
	UnlinkRepositoryFromProject                                 *UnlinkRepositoryFromProjectPayload                                 "json:\"unlinkRepositoryFromProject\" graphql:\"unlinkRepositoryFromProject\""
	UnlockLockable                                              *UnlockLockablePayload                                              "json:\"unlockLockable\" graphql:\"unlockLockable\""
	UnmarkDiscussionCommentAsAnswer                             *UnmarkDiscussionCommentAsAnswerPayload                             "json:\"unmarkDiscussionCommentAsAnswer\" graphql:\"unmarkDiscussionCommentAsAnswer\""
	UnmarkFileAsViewed                                          *UnmarkFileAsViewedPayload                                          "json:\"unmarkFileAsViewed\" graphql:\"unmarkFileAsViewed\""
	UnmarkIssueAsDuplicate                                      *UnmarkIssueAsDuplicatePayload                                      "json:\"unmarkIssueAsDuplicate\" graphql:\"unmarkIssueAsDuplicate\""
	UnminimizeComment                                           *UnminimizeCommentPayload                                           "json:\"unminimizeComment\" graphql:\"unminimizeComment\""
	UnpinIssue                                                  *UnpinIssuePayload                                                  "json:\"unpinIssue\" graphql:\"unpinIssue\""
	UnresolveReviewThread                                       *UnresolveReviewThreadPayload                                       "json:\"unresolveReviewThread\" graphql:\"unresolveReviewThread\""
	UpdateBranchProtectionRule                                  *UpdateBranchProtectionRulePayload                                  "json:\"updateBranchProtectionRule\" graphql:\"updateBranchProtectionRule\""
	UpdateCheckRun                                              *UpdateCheckRunPayload                                              "json:\"updateCheckRun\" graphql:\"updateCheckRun\""
	UpdateCheckSuitePreferences                                 *UpdateCheckSuitePreferencesPayload                                 "json:\"updateCheckSuitePreferences\" graphql:\"updateCheckSuitePreferences\""
	UpdateDiscussion                                            *UpdateDiscussionPayload                                            "json:\"updateDiscussion\" graphql:\"updateDiscussion\""
	UpdateDiscussionComment                                     *UpdateDiscussionCommentPayload                                     "json:\"updateDiscussionComment\" graphql:\"updateDiscussionComment\""
	UpdateEnterpriseAdministratorRole                           *UpdateEnterpriseAdministratorRolePayload                           "json:\"updateEnterpriseAdministratorRole\" graphql:\"updateEnterpriseAdministratorRole\""
	UpdateEnterpriseAllowPrivateRepositoryForkingSetting        *UpdateEnterpriseAllowPrivateRepositoryForkingSettingPayload        "json:\"updateEnterpriseAllowPrivateRepositoryForkingSetting\" graphql:\"updateEnterpriseAllowPrivateRepositoryForkingSetting\""
	UpdateEnterpriseDefaultRepositoryPermissionSetting          *UpdateEnterpriseDefaultRepositoryPermissionSettingPayload          "json:\"updateEnterpriseDefaultRepositoryPermissionSetting\" graphql:\"updateEnterpriseDefaultRepositoryPermissionSetting\""
	UpdateEnterpriseMembersCanChangeRepositoryVisibilitySetting *UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingPayload "json:\"updateEnterpriseMembersCanChangeRepositoryVisibilitySetting\" graphql:\"updateEnterpriseMembersCanChangeRepositoryVisibilitySetting\""
	UpdateEnterpriseMembersCanCreateRepositoriesSetting         *UpdateEnterpriseMembersCanCreateRepositoriesSettingPayload         "json:\"updateEnterpriseMembersCanCreateRepositoriesSetting\" graphql:\"updateEnterpriseMembersCanCreateRepositoriesSetting\""
	UpdateEnterpriseMembersCanDeleteIssuesSetting               *UpdateEnterpriseMembersCanDeleteIssuesSettingPayload               "json:\"updateEnterpriseMembersCanDeleteIssuesSetting\" graphql:\"updateEnterpriseMembersCanDeleteIssuesSetting\""
	UpdateEnterpriseMembersCanDeleteRepositoriesSetting         *UpdateEnterpriseMembersCanDeleteRepositoriesSettingPayload         "json:\"updateEnterpriseMembersCanDeleteRepositoriesSetting\" graphql:\"updateEnterpriseMembersCanDeleteRepositoriesSetting\""
	UpdateEnterpriseMembersCanInviteCollaboratorsSetting        *UpdateEnterpriseMembersCanInviteCollaboratorsSettingPayload        "json:\"updateEnterpriseMembersCanInviteCollaboratorsSetting\" graphql:\"updateEnterpriseMembersCanInviteCollaboratorsSetting\""
	UpdateEnterpriseMembersCanMakePurchasesSetting              *UpdateEnterpriseMembersCanMakePurchasesSettingPayload              "json:\"updateEnterpriseMembersCanMakePurchasesSetting\" graphql:\"updateEnterpriseMembersCanMakePurchasesSetting\""
	UpdateEnterpriseMembersCanUpdateProtectedBranchesSetting    *UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingPayload    "json:\"updateEnterpriseMembersCanUpdateProtectedBranchesSetting\" graphql:\"updateEnterpriseMembersCanUpdateProtectedBranchesSetting\""
	UpdateEnterpriseMembersCanViewDependencyInsightsSetting     *UpdateEnterpriseMembersCanViewDependencyInsightsSettingPayload     "json:\"updateEnterpriseMembersCanViewDependencyInsightsSetting\" graphql:\"updateEnterpriseMembersCanViewDependencyInsightsSetting\""
	UpdateEnterpriseOrganizationProjectsSetting                 *UpdateEnterpriseOrganizationProjectsSettingPayload                 "json:\"updateEnterpriseOrganizationProjectsSetting\" graphql:\"updateEnterpriseOrganizationProjectsSetting\""
	UpdateEnterpriseProfile                                     *UpdateEnterpriseProfilePayload                                     "json:\"updateEnterpriseProfile\" graphql:\"updateEnterpriseProfile\""
	UpdateEnterpriseRepositoryProjectsSetting                   *UpdateEnterpriseRepositoryProjectsSettingPayload                   "json:\"updateEnterpriseRepositoryProjectsSetting\" graphql:\"updateEnterpriseRepositoryProjectsSetting\""
	UpdateEnterpriseTeamDiscussionsSetting                      *UpdateEnterpriseTeamDiscussionsSettingPayload                      "json:\"updateEnterpriseTeamDiscussionsSetting\" graphql:\"updateEnterpriseTeamDiscussionsSetting\""
	UpdateEnterpriseTwoFactorAuthenticationRequiredSetting      *UpdateEnterpriseTwoFactorAuthenticationRequiredSettingPayload      "json:\"updateEnterpriseTwoFactorAuthenticationRequiredSetting\" graphql:\"updateEnterpriseTwoFactorAuthenticationRequiredSetting\""
	UpdateEnvironment                                           *UpdateEnvironmentPayload                                           "json:\"updateEnvironment\" graphql:\"updateEnvironment\""
	UpdateIPAllowListEnabledSetting                             *UpdateIPAllowListEnabledSettingPayload                             "json:\"updateIpAllowListEnabledSetting\" graphql:\"updateIpAllowListEnabledSetting\""
	UpdateIPAllowListEntry                                      *UpdateIPAllowListEntryPayload                                      "json:\"updateIpAllowListEntry\" graphql:\"updateIpAllowListEntry\""
	UpdateIPAllowListForInstalledAppsEnabledSetting             *UpdateIPAllowListForInstalledAppsEnabledSettingPayload             "json:\"updateIpAllowListForInstalledAppsEnabledSetting\" graphql:\"updateIpAllowListForInstalledAppsEnabledSetting\""
	UpdateIssue                                                 *UpdateIssuePayload                                                 "json:\"updateIssue\" graphql:\"updateIssue\""
	UpdateIssueComment                                          *UpdateIssueCommentPayload                                          "json:\"updateIssueComment\" graphql:\"updateIssueComment\""
	UpdateNotificationRestrictionSetting                        *UpdateNotificationRestrictionSettingPayload                        "json:\"updateNotificationRestrictionSetting\" graphql:\"updateNotificationRestrictionSetting\""
	UpdateProject                                               *UpdateProjectPayload                                               "json:\"updateProject\" graphql:\"updateProject\""
	UpdateProjectCard                                           *UpdateProjectCardPayload                                           "json:\"updateProjectCard\" graphql:\"updateProjectCard\""
	UpdateProjectColumn                                         *UpdateProjectColumnPayload                                         "json:\"updateProjectColumn\" graphql:\"updateProjectColumn\""
	UpdatePullRequest                                           *UpdatePullRequestPayload                                           "json:\"updatePullRequest\" graphql:\"updatePullRequest\""
	UpdatePullRequestReview                                     *UpdatePullRequestReviewPayload                                     "json:\"updatePullRequestReview\" graphql:\"updatePullRequestReview\""
	UpdatePullRequestReviewComment                              *UpdatePullRequestReviewCommentPayload                              "json:\"updatePullRequestReviewComment\" graphql:\"updatePullRequestReviewComment\""
	UpdateRef                                                   *UpdateRefPayload                                                   "json:\"updateRef\" graphql:\"updateRef\""
	UpdateRepository                                            *UpdateRepositoryPayload                                            "json:\"updateRepository\" graphql:\"updateRepository\""
	UpdateSubscription                                          *UpdateSubscriptionPayload                                          "json:\"updateSubscription\" graphql:\"updateSubscription\""
	UpdateTeamDiscussion                                        *UpdateTeamDiscussionPayload                                        "json:\"updateTeamDiscussion\" graphql:\"updateTeamDiscussion\""
	UpdateTeamDiscussionComment                                 *UpdateTeamDiscussionCommentPayload                                 "json:\"updateTeamDiscussionComment\" graphql:\"updateTeamDiscussionComment\""
	UpdateTopics                                                *UpdateTopicsPayload                                                "json:\"updateTopics\" graphql:\"updateTopics\""
	VerifyVerifiableDomain                                      *VerifyVerifiableDomainPayload                                      "json:\"verifyVerifiableDomain\" graphql:\"verifyVerifiableDomain\""
}
type GetTags struct {
	Repository *struct {
		Refs *struct {
			Nodes []*struct {
				Name string "json:\"name\" graphql:\"name\""
			} "json:\"nodes\" graphql:\"nodes\""
		} "json:\"refs\" graphql:\"refs\""
	} "json:\"repository\" graphql:\"repository\""
}

const GetTagsDocument = `query GetTags ($owner: String!, $name: String!, $first: Int = 30) {
	repository(owner: $owner, name: $name) {
		refs(refPrefix: "refs/tags/", first: $first, orderBy: {field:TAG_COMMIT_DATE,direction:DESC}) {
			nodes {
				name
			}
		}
	}
}
`

func (c *Client) GetTags(ctx context.Context, owner string, name string, first *int, interceptors ...clientv2.RequestInterceptor) (*GetTags, error) {
	vars := map[string]interface{}{
		"owner": owner,
		"name":  name,
		"first": first,
	}

	var res GetTags
	if err := c.Client.Post(ctx, "GetTags", GetTagsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
