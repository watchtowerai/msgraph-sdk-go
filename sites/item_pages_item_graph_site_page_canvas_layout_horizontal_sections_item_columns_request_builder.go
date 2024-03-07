package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetQueryParameters get columns from sites
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetQueryParameters struct {
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
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHorizontalSectionColumnId provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) ByHorizontalSectionColumnId(horizontalSectionColumnId string)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if horizontalSectionColumnId != "" {
        urlTplParams["horizontalSectionColumn%2Did"] = horizontalSectionColumnId
    }
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderInternal instantiates a new ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) {
    m := &ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder instantiates a new ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder and sets the default values.
func NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) Count()(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilder) {
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get columns from sites
// returns a HorizontalSectionColumnCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnCollectionResponseable), nil
}
// Post create new navigation property to columns for sites
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable), nil
}
// ToGetRequestInformation get columns from sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to columns for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder when successful
func (m *ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder) {
    return NewItemPagesItemGraphSitePageCanvasLayoutHorizontalSectionsItemColumnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
