package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
type DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal instantiates a new GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    m := &DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceManagementIntentPerSettingContributingProfiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder instantiates a new GetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder and sets the default values.
func NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesPostRequestBodyable, requestConfiguration *DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getDeviceManagementIntentPerSettingContributingProfiles
func (m *DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesPostRequestBodyable, requestConfiguration *DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}