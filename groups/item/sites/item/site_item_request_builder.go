package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i19d3b905956d0d5ab3e405e912ce5512047d11c0368cfb7b3532e4e86e5a07af "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores"
    i23ed04a56449d98fcd568705a0cf04a1b544ee9821f049f84ae0c25b1d8a1f76 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/items"
    i433a1560954e7a6c21be030d3739b5c8e7a3a62b25192a404e6284478dbdf361 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/columns"
    i4530be9f2649e2da7651adc83045d48e03ce2d02fe5bbfac583977b13834180c "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists"
    i5ff933dfab28b9bc528e49bf8331befadb5004877c8387d60f5e76079d8ba1a6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/externalcolumns"
    i7bf3dbc49a37a69a92d9c0b2f5894dbd2387cd5cf2c265de0a9f1307dedf7007 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/onenote"
    i9619fa512c16d940bb8a99f87dcbb11e238865522dd9f8aa27e8eb6839d210c8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/sites"
    iabe6a78fc3f592fa955e512cf883dcc83f9a1c269265fa5904300f3f3c66d279 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstore"
    ibe39270ba4e624ff4fe340aa2bd6d3631b4dd3b549fc6c5dbafcf7756c2cae95 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/permissions"
    ida62d1e3e0c3a4155187312adb160423df6f1b84474a7c37dac6b38ee0d57c30 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes"
    ie2e0fba360015570188988834e1d1616938882b51276657b518960b0529c8395 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/drive"
    ie977d634115f6edbcba02945f5f558e1f1060c3078fc1043cfc6a70da5598048 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/analytics"
    if423f3b786c94937b52a1c87021c6cf25ba0bba2b0405608a8d1629f108d2664 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/drives"
    i20d71e187e2869b1bf133549cf92b03dbdd4fcd7a6312d7e4cfab185397ec504 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/drives/item"
    i3d15a6f13c74ebb900c3f4d7e8b8cf4dc95de03aed0a0c4b8a2a4e6ed582e97e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/columns/item"
    i3eb82c303a8d1465dafb4c1bc15b34ff40013f94aac87f058c363a7aba92e4e8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/externalcolumns/item"
    i542a41cec057b321ef1e41187fa571edf863979b732f2c19975560ced8e76f8b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/permissions/item"
    ia63bedd538371b0186a5b6eee2c0d77c2ae0f708dcde2fcdeeebb39fe26bd8b7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/termstores/item"
    iba2aa0555939eae7c495b9747ca02d399d5a9d9768108b185661cc038711fdbc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/items/item"
    ic3cb2a24c56a0f014badb537b2e05366face417e10e73eda79e507ce267e5183 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/contenttypes/item"
    icb7ef07a689e39fb1ed0e6596004ecf47a3815ad87f19c96bff8fb7ad90304b9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/sites/item"
    iea900cf48c5d3da0d6a702e72cfb722443dcaaa70a27924a2c1a8b756a55448d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/sites/item/lists/item"
)

// SiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type SiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SiteItemRequestBuilderDeleteOptions options for Delete
type SiteItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SiteItemRequestBuilderGetOptions options for Get
type SiteItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SiteItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type SiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SiteItemRequestBuilderPatchOptions options for Patch
type SiteItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Analytics the analytics property
func (m *SiteItemRequestBuilder) Analytics()(*ie977d634115f6edbcba02945f5f558e1f1060c3078fc1043cfc6a70da5598048.AnalyticsRequestBuilder) {
    return ie977d634115f6edbcba02945f5f558e1f1060c3078fc1043cfc6a70da5598048.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *SiteItemRequestBuilder) Columns()(*i433a1560954e7a6c21be030d3739b5c8e7a3a62b25192a404e6284478dbdf361.ColumnsRequestBuilder) {
    return i433a1560954e7a6c21be030d3739b5c8e7a3a62b25192a404e6284478dbdf361.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.columns.item collection
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*i3d15a6f13c74ebb900c3f4d7e8b8cf4dc95de03aed0a0c4b8a2a4e6ed582e97e.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i3d15a6f13c74ebb900c3f4d7e8b8cf4dc95de03aed0a0c4b8a2a4e6ed582e97e.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *SiteItemRequestBuilder) ContentTypes()(*ida62d1e3e0c3a4155187312adb160423df6f1b84474a7c37dac6b38ee0d57c30.ContentTypesRequestBuilder) {
    return ida62d1e3e0c3a4155187312adb160423df6f1b84474a7c37dac6b38ee0d57c30.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.contentTypes.item collection
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*ic3cb2a24c56a0f014badb537b2e05366face417e10e73eda79e507ce267e5183.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return ic3cb2a24c56a0f014badb537b2e05366face417e10e73eda79e507ce267e5183.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sites for groups
func (m *SiteItemRequestBuilder) CreateDeleteRequestInformation(options *SiteItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) CreateGetRequestInformation(options *SiteItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sites in groups
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(options *SiteItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property sites for groups
func (m *SiteItemRequestBuilder) Delete(options *SiteItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive the drive property
func (m *SiteItemRequestBuilder) Drive()(*ie2e0fba360015570188988834e1d1616938882b51276657b518960b0529c8395.DriveRequestBuilder) {
    return ie2e0fba360015570188988834e1d1616938882b51276657b518960b0529c8395.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *SiteItemRequestBuilder) Drives()(*if423f3b786c94937b52a1c87021c6cf25ba0bba2b0405608a8d1629f108d2664.DrivesRequestBuilder) {
    return if423f3b786c94937b52a1c87021c6cf25ba0bba2b0405608a8d1629f108d2664.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.drives.item collection
func (m *SiteItemRequestBuilder) DrivesById(id string)(*i20d71e187e2869b1bf133549cf92b03dbdd4fcd7a6312d7e4cfab185397ec504.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i20d71e187e2869b1bf133549cf92b03dbdd4fcd7a6312d7e4cfab185397ec504.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns the externalColumns property
func (m *SiteItemRequestBuilder) ExternalColumns()(*i5ff933dfab28b9bc528e49bf8331befadb5004877c8387d60f5e76079d8ba1a6.ExternalColumnsRequestBuilder) {
    return i5ff933dfab28b9bc528e49bf8331befadb5004877c8387d60f5e76079d8ba1a6.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.externalColumns.item collection
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*i3eb82c303a8d1465dafb4c1bc15b34ff40013f94aac87f058c363a7aba92e4e8.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i3eb82c303a8d1465dafb4c1bc15b34ff40013f94aac87f058c363a7aba92e4e8.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) Get(options *SiteItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// Items the items property
func (m *SiteItemRequestBuilder) Items()(*i23ed04a56449d98fcd568705a0cf04a1b544ee9821f049f84ae0c25b1d8a1f76.ItemsRequestBuilder) {
    return i23ed04a56449d98fcd568705a0cf04a1b544ee9821f049f84ae0c25b1d8a1f76.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.items.item collection
func (m *SiteItemRequestBuilder) ItemsById(id string)(*iba2aa0555939eae7c495b9747ca02d399d5a9d9768108b185661cc038711fdbc.BaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return iba2aa0555939eae7c495b9747ca02d399d5a9d9768108b185661cc038711fdbc.NewBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists the lists property
func (m *SiteItemRequestBuilder) Lists()(*i4530be9f2649e2da7651adc83045d48e03ce2d02fe5bbfac583977b13834180c.ListsRequestBuilder) {
    return i4530be9f2649e2da7651adc83045d48e03ce2d02fe5bbfac583977b13834180c.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.lists.item collection
func (m *SiteItemRequestBuilder) ListsById(id string)(*iea900cf48c5d3da0d6a702e72cfb722443dcaaa70a27924a2c1a8b756a55448d.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return iea900cf48c5d3da0d6a702e72cfb722443dcaaa70a27924a2c1a8b756a55448d.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *SiteItemRequestBuilder) Onenote()(*i7bf3dbc49a37a69a92d9c0b2f5894dbd2387cd5cf2c265de0a9f1307dedf7007.OnenoteRequestBuilder) {
    return i7bf3dbc49a37a69a92d9c0b2f5894dbd2387cd5cf2c265de0a9f1307dedf7007.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sites in groups
func (m *SiteItemRequestBuilder) Patch(options *SiteItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *SiteItemRequestBuilder) Permissions()(*ibe39270ba4e624ff4fe340aa2bd6d3631b4dd3b549fc6c5dbafcf7756c2cae95.PermissionsRequestBuilder) {
    return ibe39270ba4e624ff4fe340aa2bd6d3631b4dd3b549fc6c5dbafcf7756c2cae95.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.permissions.item collection
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*i542a41cec057b321ef1e41187fa571edf863979b732f2c19975560ced8e76f8b.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i542a41cec057b321ef1e41187fa571edf863979b732f2c19975560ced8e76f8b.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *SiteItemRequestBuilder) Sites()(*i9619fa512c16d940bb8a99f87dcbb11e238865522dd9f8aa27e8eb6839d210c8.SitesRequestBuilder) {
    return i9619fa512c16d940bb8a99f87dcbb11e238865522dd9f8aa27e8eb6839d210c8.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.sites.item collection
func (m *SiteItemRequestBuilder) SitesById(id string)(*icb7ef07a689e39fb1ed0e6596004ecf47a3815ad87f19c96bff8fb7ad90304b9.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return icb7ef07a689e39fb1ed0e6596004ecf47a3815ad87f19c96bff8fb7ad90304b9.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore the termStore property
func (m *SiteItemRequestBuilder) TermStore()(*iabe6a78fc3f592fa955e512cf883dcc83f9a1c269265fa5904300f3f3c66d279.TermStoreRequestBuilder) {
    return iabe6a78fc3f592fa955e512cf883dcc83f9a1c269265fa5904300f3f3c66d279.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStores the termStores property
func (m *SiteItemRequestBuilder) TermStores()(*i19d3b905956d0d5ab3e405e912ce5512047d11c0368cfb7b3532e4e86e5a07af.TermStoresRequestBuilder) {
    return i19d3b905956d0d5ab3e405e912ce5512047d11c0368cfb7b3532e4e86e5a07af.NewTermStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStoresById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.sites.item.termStores.item collection
func (m *SiteItemRequestBuilder) TermStoresById(id string)(*ia63bedd538371b0186a5b6eee2c0d77c2ae0f708dcde2fcdeeebb39fe26bd8b7.StoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["store%2Did"] = id
    }
    return ia63bedd538371b0186a5b6eee2c0d77c2ae0f708dcde2fcdeeebb39fe26bd8b7.NewStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}