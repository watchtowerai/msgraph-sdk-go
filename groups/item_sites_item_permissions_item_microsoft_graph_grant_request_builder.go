package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder provides operations to call the grant method.
type ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderInternal instantiates a new MicrosoftGraphGrantRequestBuilder and sets the default values.
func NewItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder) {
    m := &ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/permissions/{permission%2Did}/microsoft.graph.grant";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder instantiates a new MicrosoftGraphGrantRequestBuilder and sets the default values.
func NewItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post grant users access to a link represented by a [permission][].
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/permission-grant?view=graph-rest-1.0
func (m *ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder) Post(ctx context.Context, body ItemSitesItemPermissionsItemMicrosoftGraphGrantGrantPostRequestBodyable, requestConfiguration *ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderPostRequestConfiguration)(ItemSitesItemPermissionsItemMicrosoftGraphGrantGrantResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemSitesItemPermissionsItemMicrosoftGraphGrantGrantResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemPermissionsItemMicrosoftGraphGrantGrantResponseable), nil
}
// ToPostRequestInformation grant users access to a link represented by a [permission][].
func (m *ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemPermissionsItemMicrosoftGraphGrantGrantPostRequestBodyable, requestConfiguration *ItemSitesItemPermissionsItemMicrosoftGraphGrantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}