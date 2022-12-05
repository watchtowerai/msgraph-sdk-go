package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemChatsItemSendActivityNotificationPostRequestBodyable 
type UsersItemChatsItemSendActivityNotificationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityType()(*string)
    GetChainId()(*int64)
    GetPreviewText()(ItemBodyable)
    GetRecipient()(TeamworkNotificationRecipientable)
    GetTemplateParameters()([]KeyValuePairable)
    GetTopic()(TeamworkActivityTopicable)
    SetActivityType(value *string)()
    SetChainId(value *int64)()
    SetPreviewText(value ItemBodyable)()
    SetRecipient(value TeamworkNotificationRecipientable)()
    SetTemplateParameters(value []KeyValuePairable)()
    SetTopic(value TeamworkActivityTopicable)()
}