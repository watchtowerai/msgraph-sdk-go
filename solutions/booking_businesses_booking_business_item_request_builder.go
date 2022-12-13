package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BookingBusinessesBookingBusinessItemRequestBuilder provides operations to manage the bookingBusinesses property of the microsoft.graph.solutionsRoot entity.
type BookingBusinessesBookingBusinessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters get bookingBusinesses from solutions
type BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookingBusinessesBookingBusinessItemRequestBuilderGetQueryParameters
}
// BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Appointments provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Appointments()(*BookingBusinessesItemAppointmentsRequestBuilder) {
    return NewBookingBusinessesItemAppointmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppointmentsById provides operations to manage the appointments property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) AppointmentsById(id string)(*BookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewBookingBusinessesItemAppointmentsBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CalendarView()(*BookingBusinessesItemCalendarViewRequestBuilder) {
    return NewBookingBusinessesItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CalendarViewById(id string)(*BookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingAppointment%2Did"] = id
    }
    return NewBookingBusinessesItemCalendarViewBookingAppointmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewBookingBusinessesBookingBusinessItemRequestBuilderInternal instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessesBookingBusinessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesBookingBusinessItemRequestBuilder) {
    m := &BookingBusinessesBookingBusinessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookingBusinessesBookingBusinessItemRequestBuilder instantiates a new BookingBusinessItemRequestBuilder and sets the default values.
func NewBookingBusinessesBookingBusinessItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookingBusinessesBookingBusinessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookingBusinessesBookingBusinessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get bookingBusinesses from solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Customers provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Customers()(*BookingBusinessesItemCustomersRequestBuilder) {
    return NewBookingBusinessesItemCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomersById(id string)(*BookingBusinessesItemCustomersBookingCustomerBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomerBase%2Did"] = id
    }
    return NewBookingBusinessesItemCustomersBookingCustomerBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomQuestions provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestions()(*BookingBusinessesItemCustomQuestionsRequestBuilder) {
    return NewBookingBusinessesItemCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomQuestionsById provides operations to manage the customQuestions property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) CustomQuestionsById(id string)(*BookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCustomQuestion%2Did"] = id
    }
    return NewBookingBusinessesItemCustomQuestionsBookingCustomQuestionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property bookingBusinesses for solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get bookingBusinesses from solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable), nil
}
// GetStaffAvailability provides operations to call the getStaffAvailability method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) GetStaffAvailability()(*BookingBusinessesItemGetStaffAvailabilityRequestBuilder) {
    return NewBookingBusinessesItemGetStaffAvailabilityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bookingBusinesses in solutions
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, requestConfiguration *BookingBusinessesBookingBusinessItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingBusinessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingBusinessable), nil
}
// Publish provides operations to call the publish method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Publish()(*BookingBusinessesItemPublishRequestBuilder) {
    return NewBookingBusinessesItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Services provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Services()(*BookingBusinessesItemServicesRequestBuilder) {
    return NewBookingBusinessesItemServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicesById provides operations to manage the services property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) ServicesById(id string)(*BookingBusinessesItemServicesBookingServiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingService%2Did"] = id
    }
    return NewBookingBusinessesItemServicesBookingServiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// StaffMembers provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) StaffMembers()(*BookingBusinessesItemStaffMembersRequestBuilder) {
    return NewBookingBusinessesItemStaffMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StaffMembersById provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) StaffMembersById(id string)(*BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingStaffMemberBase%2Did"] = id
    }
    return NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *BookingBusinessesBookingBusinessItemRequestBuilder) Unpublish()(*BookingBusinessesItemUnpublishRequestBuilder) {
    return NewBookingBusinessesItemUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
