package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody provides operations to call the validateProperties method.
type IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The entityType property
    entityType *string
    // The mailNickname property
    mailNickname *string
    // The onBehalfOfUserId property
    onBehalfOfUserId *string
}
// NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody instantiates a new IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody and sets the default values.
func NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) {
    m := &IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEntityType gets the entityType property value. The entityType property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetEntityType()(*string) {
    return m.entityType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["entityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["onBehalfOfUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOfUserId(val)
        }
        return nil
    }
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetOnBehalfOfUserId gets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) GetOnBehalfOfUserId()(*string) {
    return m.onBehalfOfUserId
}
// Serialize serializes information the current object
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEntityType sets the entityType property value. The entityType property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) SetEntityType(value *string)() {
    m.entityType = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetOnBehalfOfUserId sets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *IdentityGovernanceEntitlementManagementConnectedOrganizationsItemExternalSponsorsValidatePropertiesPostRequestBody) SetOnBehalfOfUserId(value *string)() {
    m.onBehalfOfUserId = value
}