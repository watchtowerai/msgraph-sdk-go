package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewScheduleDefinition 
type AccessReviewScheduleDefinition struct {
    Entity
    // Defines the list of additional users or group members to be notified of the access review progress.
    additionalNotificationRecipients []AccessReviewNotificationRecipientItem;
    // User who created this review. Read-only.
    createdBy *UserIdentity;
    // Timestamp when the access review series was created. Supports $select and $orderBy. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description provided by review creators to provide more context of the review to admins. Supports $select.
    descriptionForAdmins *string;
    // Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
    descriptionForReviewers *string;
    // Name of the access review series. Supports $select and $orderBy. Required on create.
    displayName *string;
    // This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select.
    fallbackReviewers []AccessReviewReviewerScope;
    // This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
    instanceEnumerationScope *AccessReviewScope;
    // If the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
    instances []AccessReviewInstance;
    // Timestamp when the access review series was last modified. Supports $select. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
    reviewers []AccessReviewReviewerScope;
    // Defines the entities whose access is reviewed.  For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
    scope *AccessReviewScope;
    // The settings for an access review series, see type definition below. Supports $select. Required on create.
    settings *AccessReviewScheduleSettings;
    // This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
    status *string;
}
// NewAccessReviewScheduleDefinition instantiates a new accessReviewScheduleDefinition and sets the default values.
func NewAccessReviewScheduleDefinition()(*AccessReviewScheduleDefinition) {
    m := &AccessReviewScheduleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdditionalNotificationRecipients gets the additionalNotificationRecipients property value. Defines the list of additional users or group members to be notified of the access review progress.
func (m *AccessReviewScheduleDefinition) GetAdditionalNotificationRecipients()([]AccessReviewNotificationRecipientItem) {
    if m == nil {
        return nil
    } else {
        return m.additionalNotificationRecipients
    }
}
// GetCreatedBy gets the createdBy property value. User who created this review. Read-only.
func (m *AccessReviewScheduleDefinition) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp when the access review series was created. Supports $select and $orderBy. Read-only.
func (m *AccessReviewScheduleDefinition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescriptionForAdmins gets the descriptionForAdmins property value. Description provided by review creators to provide more context of the review to admins. Supports $select.
func (m *AccessReviewScheduleDefinition) GetDescriptionForAdmins()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForAdmins
    }
}
// GetDescriptionForReviewers gets the descriptionForReviewers property value. Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
func (m *AccessReviewScheduleDefinition) GetDescriptionForReviewers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForReviewers
    }
}
// GetDisplayName gets the displayName property value. Name of the access review series. Supports $select and $orderBy. Required on create.
func (m *AccessReviewScheduleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFallbackReviewers gets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select.
func (m *AccessReviewScheduleDefinition) GetFallbackReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.fallbackReviewers
    }
}
// GetInstanceEnumerationScope gets the instanceEnumerationScope property value. This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) GetInstanceEnumerationScope()(*AccessReviewScope) {
    if m == nil {
        return nil
    } else {
        return m.instanceEnumerationScope
    }
}
// GetInstances gets the instances property value. If the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewScheduleDefinition) GetInstances()([]AccessReviewInstance) {
    if m == nil {
        return nil
    } else {
        return m.instances
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp when the access review series was last modified. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetReviewers gets the reviewers property value. This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// GetScope gets the scope property value. Defines the entities whose access is reviewed.  For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) GetScope()(*AccessReviewScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetSettings gets the settings property value. The settings for an access review series, see type definition below. Supports $select. Required on create.
func (m *AccessReviewScheduleDefinition) GetSettings()(*AccessReviewScheduleSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetStatus gets the status property value. This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewScheduleDefinition) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewScheduleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalNotificationRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewNotificationRecipientItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewNotificationRecipientItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewNotificationRecipientItem))
            }
            m.SetAdditionalNotificationRecipients(res)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*UserIdentity))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["descriptionForAdmins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForAdmins(val)
        }
        return nil
    }
    res["descriptionForReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescriptionForReviewers(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["fallbackReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewReviewerScope))
            }
            m.SetFallbackReviewers(res)
        }
        return nil
    }
    res["instanceEnumerationScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScope() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceEnumerationScope(val.(*AccessReviewScope))
        }
        return nil
    }
    res["instances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewInstance))
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewReviewerScope))
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScope() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*AccessReviewScope))
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScheduleSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*AccessReviewScheduleSettings))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *AccessReviewScheduleDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewScheduleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalNotificationRecipients()))
        for i, v := range m.GetAdditionalNotificationRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("additionalNotificationRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForAdmins", m.GetDescriptionForAdmins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForReviewers", m.GetDescriptionForReviewers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFallbackReviewers()))
        for i, v := range m.GetFallbackReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instanceEnumerationScope", m.GetInstanceEnumerationScope())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalNotificationRecipients sets the additionalNotificationRecipients property value. Defines the list of additional users or group members to be notified of the access review progress.
func (m *AccessReviewScheduleDefinition) SetAdditionalNotificationRecipients(value []AccessReviewNotificationRecipientItem)() {
    if m != nil {
        m.additionalNotificationRecipients = value
    }
}
// SetCreatedBy sets the createdBy property value. User who created this review. Read-only.
func (m *AccessReviewScheduleDefinition) SetCreatedBy(value *UserIdentity)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp when the access review series was created. Supports $select and $orderBy. Read-only.
func (m *AccessReviewScheduleDefinition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescriptionForAdmins sets the descriptionForAdmins property value. Description provided by review creators to provide more context of the review to admins. Supports $select.
func (m *AccessReviewScheduleDefinition) SetDescriptionForAdmins(value *string)() {
    if m != nil {
        m.descriptionForAdmins = value
    }
}
// SetDescriptionForReviewers sets the descriptionForReviewers property value. Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
func (m *AccessReviewScheduleDefinition) SetDescriptionForReviewers(value *string)() {
    if m != nil {
        m.descriptionForReviewers = value
    }
}
// SetDisplayName sets the displayName property value. Name of the access review series. Supports $select and $orderBy. Required on create.
func (m *AccessReviewScheduleDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFallbackReviewers sets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select.
func (m *AccessReviewScheduleDefinition) SetFallbackReviewers(value []AccessReviewReviewerScope)() {
    if m != nil {
        m.fallbackReviewers = value
    }
}
// SetInstanceEnumerationScope sets the instanceEnumerationScope property value. This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) SetInstanceEnumerationScope(value *AccessReviewScope)() {
    if m != nil {
        m.instanceEnumerationScope = value
    }
}
// SetInstances sets the instances property value. If the accessReviewScheduleDefinition is a recurring access review, instances represent each recurrence. A review that does not recur will have exactly one instance. Instances also represent each unique resource under review in the accessReviewScheduleDefinition. If a review has multiple resources and multiple instances, each resource will have a unique instance for each recurrence.
func (m *AccessReviewScheduleDefinition) SetInstances(value []AccessReviewInstance)() {
    if m != nil {
        m.instances = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp when the access review series was last modified. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetReviewers sets the reviewers property value. This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) SetReviewers(value []AccessReviewReviewerScope)() {
    if m != nil {
        m.reviewers = value
    }
}
// SetScope sets the scope property value. Defines the entities whose access is reviewed.  For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) SetScope(value *AccessReviewScope)() {
    if m != nil {
        m.scope = value
    }
}
// SetSettings sets the settings property value. The settings for an access review series, see type definition below. Supports $select. Required on create.
func (m *AccessReviewScheduleDefinition) SetSettings(value *AccessReviewScheduleSettings)() {
    if m != nil {
        m.settings = value
    }
}
// SetStatus sets the status property value. This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewScheduleDefinition) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
