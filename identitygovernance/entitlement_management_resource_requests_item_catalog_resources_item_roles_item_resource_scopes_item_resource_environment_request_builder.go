package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder provides operations to manage the environment property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
type EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetQueryParameters
}
// NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal instantiates a new EnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) {
    m := &EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resourceRequests/{accessPackageResourceRequest%2Did}/catalog/resources/{accessPackageResource%2Did}/roles/{accessPackageResourceRole%2Did}/resource/scopes/{accessPackageResourceScope%2Did}/resource/environment{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder instantiates a new EnvironmentRequestBuilder and sets the default values.
func NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation contains the environment information for the resource. This can be set using either the @odata.bind annotation or the environment's originId.Supports $expand.
func (m *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder) {
    return NewEntitlementManagementResourceRequestsItemCatalogResourcesItemRolesItemResourceScopesItemResourceEnvironmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
