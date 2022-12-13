package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
type ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters the managed app registrations.
type ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters
}
// ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliedPolicies provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) AppliedPolicies()(*ManagedAppRegistrationsItemAppliedPoliciesRequestBuilder) {
    return NewManagedAppRegistrationsItemAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) AppliedPoliciesById(id string)(*ManagedAppRegistrationsItemAppliedPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return NewManagedAppRegistrationsItemAppliedPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    m := &ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property managedAppRegistrations for deviceAppManagement
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the managed app registrations.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
// IntendedPolicies provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) IntendedPolicies()(*ManagedAppRegistrationsItemIntendedPoliciesRequestBuilder) {
    return NewManagedAppRegistrationsItemIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) IntendedPoliciesById(id string)(*ManagedAppRegistrationsItemIntendedPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return NewManagedAppRegistrationsItemIntendedPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Operations()(*ManagedAppRegistrationsItemOperationsRequestBuilder) {
    return NewManagedAppRegistrationsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.managedAppRegistration entity.
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) OperationsById(id string)(*ManagedAppRegistrationsItemOperationsManagedAppOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation%2Did"] = id
    }
    return NewManagedAppRegistrationsItemOperationsManagedAppOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedAppRegistrations in deviceAppManagement
func (m *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, requestConfiguration *ManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
