package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatsItemMarkChatReadForUserPostRequestBody provides operations to call the markChatReadForUser method.
type ChatsItemMarkChatReadForUserPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The user property
    user TeamworkUserIdentityable
}
// NewChatsItemMarkChatReadForUserPostRequestBody instantiates a new ChatsItemMarkChatReadForUserPostRequestBody and sets the default values.
func NewChatsItemMarkChatReadForUserPostRequestBody()(*ChatsItemMarkChatReadForUserPostRequestBody) {
    m := &ChatsItemMarkChatReadForUserPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatsItemMarkChatReadForUserPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatsItemMarkChatReadForUserPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatsItemMarkChatReadForUserPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatsItemMarkChatReadForUserPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatsItemMarkChatReadForUserPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetUser gets the user property value. The user property
func (m *ChatsItemMarkChatReadForUserPostRequestBody) GetUser()(TeamworkUserIdentityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *ChatsItemMarkChatReadForUserPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ChatsItemMarkChatReadForUserPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUser sets the user property value. The user property
func (m *ChatsItemMarkChatReadForUserPostRequestBody) SetUser(value TeamworkUserIdentityable)() {
    m.user = value
}