package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemReplyRequestBuilder provides operations to call the reply method.
type ItemThreadsItemPostsItemReplyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemThreadsItemPostsItemReplyRequestBuilderInternal instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemReplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemReplyRequestBuilder) {
    m := &ItemThreadsItemPostsItemReplyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/microsoft.graph.reply";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemThreadsItemPostsItemReplyRequestBuilder instantiates a new ReplyRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemReplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemReplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemReplyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation reply to a post and add a new post to the specified thread in a group conversation.  You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation.
func (m *ItemThreadsItemPostsItemReplyRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ItemThreadsItemPostsItemReplyPostRequestBodyable, requestConfiguration *ItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post reply to a post and add a new post to the specified thread in a group conversation.  You can specify both the parent conversation and thread in the request, or, you can specify just the parent thread without the parent conversation.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/post-reply?view=graph-rest-1.0
func (m *ItemThreadsItemPostsItemReplyRequestBuilder) Post(ctx context.Context, body ItemThreadsItemPostsItemReplyPostRequestBodyable, requestConfiguration *ItemThreadsItemPostsItemReplyRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
