package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OrganizationGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type OrganizationGetAvailableExtensionPropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrganizationGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrganizationGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrganizationGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewOrganizationGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationGetAvailableExtensionPropertiesRequestBuilder) {
    m := &OrganizationGetAvailableExtensionPropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/organization/microsoft.graph.getAvailableExtensionProperties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrganizationGetAvailableExtensionPropertiesRequestBuilder instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewOrganizationGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrganizationGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrganizationGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
func (m *OrganizationGetAvailableExtensionPropertiesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *OrganizationGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
func (m *OrganizationGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *OrganizationGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationGetAvailableExtensionPropertiesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOrganizationGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OrganizationGetAvailableExtensionPropertiesResponseable), nil
}