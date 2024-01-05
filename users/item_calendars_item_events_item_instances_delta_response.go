package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemEventsItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarsItemEventsItemInstancesDeltaResponse struct {
    ItemCalendarsItemEventsItemInstancesDeltaGetResponse
}
// NewItemCalendarsItemEventsItemInstancesDeltaResponse instantiates a new ItemCalendarsItemEventsItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarsItemEventsItemInstancesDeltaResponse()(*ItemCalendarsItemEventsItemInstancesDeltaResponse) {
    m := &ItemCalendarsItemEventsItemInstancesDeltaResponse{
        ItemCalendarsItemEventsItemInstancesDeltaGetResponse: *NewItemCalendarsItemEventsItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemEventsItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemEventsItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemInstancesDeltaResponse(), nil
}
// ItemCalendarsItemEventsItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarsItemEventsItemInstancesDeltaResponseable interface {
    ItemCalendarsItemEventsItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
