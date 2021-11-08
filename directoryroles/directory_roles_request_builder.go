package directoryroles

import (
    ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getavailableextensionproperties"
    ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/getbyids"
    id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1 "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/delta"
    ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a "github.com/microsoftgraph/msgraph-sdk-go/directoryroles/validateproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \directoryRoles
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type DirectoryRolesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryRolesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entities from directoryRoles
type DirectoryRolesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
}
// Options for Post
type DirectoryRolesRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new DirectoryRolesRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DirectoryRolesRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(options *DirectoryRolesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Add new entity to directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(options *DirectoryRolesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Builds and executes requests for operations under \directoryRoles\microsoft.graph.delta()
func (m *DirectoryRolesRequestBuilder) Delta()(*id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.DeltaRequestBuilder) {
    return id178ea9892b7517ea8099dc2841e11c78f3dfdc11a8a6c2abd1433ac224172e1.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entities from directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRolesRequestBuilder) Get(options *DirectoryRolesRequestBuilderGetOptions)(*DirectoryRolesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryRolesResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryRolesResponse), nil
}
func (m *DirectoryRolesRequestBuilder) GetAvailableExtensionProperties()(*ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.GetAvailableExtensionPropertiesRequestBuilder) {
    return ia4917d7d2d6bbb2d83014b7111cd4d79a23adb1d36bc4e4126b6166e1d29afa8.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryRolesRequestBuilder) GetByIds()(*ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.GetByIdsRequestBuilder) {
    return ia87332b05b41e0506ca217464387e3ed5c7d91dea72449a3a8af52c972c79b03.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to directoryRoles
// Parameters:
//  - options : Options for the request
func (m *DirectoryRolesRequestBuilder) Post(options *DirectoryRolesRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDirectoryRole() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DirectoryRole), nil
}
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.ValidatePropertiesRequestBuilder) {
    return ida2bdcfe4d6bf2bb544752c17e1ca5059ff3edc052c3080051a9c874aeb8aa1a.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
