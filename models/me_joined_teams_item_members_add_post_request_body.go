package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeJoinedTeamsItemMembersAddPostRequestBody provides operations to call the add method.
type MeJoinedTeamsItemMembersAddPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The values property
    values []ConversationMemberable
}
// NewMeJoinedTeamsItemMembersAddPostRequestBody instantiates a new MeJoinedTeamsItemMembersAddPostRequestBody and sets the default values.
func NewMeJoinedTeamsItemMembersAddPostRequestBody()(*MeJoinedTeamsItemMembersAddPostRequestBody) {
    m := &MeJoinedTeamsItemMembersAddPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeJoinedTeamsItemMembersAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeJoinedTeamsItemMembersAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeJoinedTeamsItemMembersAddPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(ConversationMemberable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetValues gets the values property value. The values property
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) GetValues()([]ConversationMemberable) {
    return m.values
}
// Serialize serializes information the current object
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("values", cast)
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
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValues sets the values property value. The values property
func (m *MeJoinedTeamsItemMembersAddPostRequestBody) SetValues(value []ConversationMemberable)() {
    m.values = value
}