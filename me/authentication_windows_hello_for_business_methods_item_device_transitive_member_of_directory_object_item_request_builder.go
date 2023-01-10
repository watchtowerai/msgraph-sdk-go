package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Application()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Device()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Group()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// User casts the previous resource to user.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) User()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
