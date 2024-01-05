package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder provides operations to count the resources in the collection.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters get the number of the resource
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters
}
// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) {
    m := &ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/children/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) {
    return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
