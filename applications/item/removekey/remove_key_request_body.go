package removekey

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoveKeyRequestBody provides operations to call the removeKey method.
type RemoveKeyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The keyId property
    keyId *string;
    // The proof property
    proof *string;
}
// NewRemoveKeyRequestBody instantiates a new removeKeyRequestBody and sets the default values.
func NewRemoveKeyRequestBody()(*RemoveKeyRequestBody) {
    m := &RemoveKeyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRemoveKeyRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoveKeyRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoveKeyRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoveKeyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoveKeyRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The keyId property
func (m *RemoveKeyRequestBody) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
// GetProof gets the proof property value. The proof property
func (m *RemoveKeyRequestBody) GetProof()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proof
    }
}
// Serialize serializes information the current object
func (m *RemoveKeyRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *RemoveKeyRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeyId sets the keyId property value. The keyId property
func (m *RemoveKeyRequestBody) SetKeyId(value *string)() {
    if m != nil {
        m.keyId = value
    }
}
// SetProof sets the proof property value. The proof property
func (m *RemoveKeyRequestBody) SetProof(value *string)() {
    if m != nil {
        m.proof = value
    }
}
