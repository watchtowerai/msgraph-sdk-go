package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody provides operations to call the createReplyAll method.
type MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
    // The Message property
    message Messageable
}
// NewMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody instantiates a new MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody and sets the default values.
func NewMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody()(*MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) {
    m := &MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(Messageable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) GetMessage()(Messageable) {
    return m.message
}
// Serialize serializes information the current object
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetMessage sets the message property value. The Message property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemCreateReplyAllPostRequestBody) SetMessage(value Messageable)() {
    m.message = value
}