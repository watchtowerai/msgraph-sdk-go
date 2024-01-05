package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementMessagesMarkUnreadResponse 
// Deprecated: This class is obsolete. Use markUnreadPostResponse instead.
type ServiceAnnouncementMessagesMarkUnreadResponse struct {
    ServiceAnnouncementMessagesMarkUnreadPostResponse
}
// NewServiceAnnouncementMessagesMarkUnreadResponse instantiates a new ServiceAnnouncementMessagesMarkUnreadResponse and sets the default values.
func NewServiceAnnouncementMessagesMarkUnreadResponse()(*ServiceAnnouncementMessagesMarkUnreadResponse) {
    m := &ServiceAnnouncementMessagesMarkUnreadResponse{
        ServiceAnnouncementMessagesMarkUnreadPostResponse: *NewServiceAnnouncementMessagesMarkUnreadPostResponse(),
    }
    return m
}
// CreateServiceAnnouncementMessagesMarkUnreadResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementMessagesMarkUnreadResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesMarkUnreadResponse(), nil
}
// ServiceAnnouncementMessagesMarkUnreadResponseable 
// Deprecated: This class is obsolete. Use markUnreadPostResponse instead.
type ServiceAnnouncementMessagesMarkUnreadResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementMessagesMarkUnreadPostResponseable
}
