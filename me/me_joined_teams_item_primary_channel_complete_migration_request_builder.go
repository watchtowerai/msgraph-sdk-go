package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewMeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) {
    m := &MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/primaryChannel/microsoft.graph.completeMigration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewMeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *MeJoinedTeamsItemPrimaryChannelCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}