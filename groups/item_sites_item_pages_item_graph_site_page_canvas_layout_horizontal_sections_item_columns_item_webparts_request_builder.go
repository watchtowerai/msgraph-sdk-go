package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters get webparts from groups
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWebPartId provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
// returns a *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ByWebPartId(webPartId string)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if webPartId != "" {
        urlTplParams["webPart%2Did"] = webPartId
    }
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsWebPartItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder instantiates a new ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Count()(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get webparts from groups
// returns a WebPartCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartCollectionResponseable), nil
}
// Post create new navigation property to webparts for groups
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable), nil
}
// ToGetRequestInformation get webparts from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to webparts for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}/webparts", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsItemWebpartsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
