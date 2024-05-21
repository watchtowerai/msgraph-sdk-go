package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder provides operations to manage the drive property of the microsoft.graph.site entity.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetQueryParameters the default drive (document library) for this site.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetQueryParameters
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderInternal instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) {
    m := &ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')/drive{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder instantiates a new ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the default drive (document library) for this site.
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Driveable), nil
}
// ToGetRequestInformation the default drive (document library) for this site.
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder when successful
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
