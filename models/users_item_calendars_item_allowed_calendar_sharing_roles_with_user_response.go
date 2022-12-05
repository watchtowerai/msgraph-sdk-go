package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse provides operations to call the allowedCalendarSharingRoles method.
type UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CalendarRoleType
}
// NewUsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse instantiates a new UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse and sets the default values.
func NewUsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse()(*UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) {
    m := &UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarRoleType, len(val))
            for i, v := range val {
                res[i] = *(v.(*CalendarRoleType))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) GetValue()([]CalendarRoleType) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", SerializeCalendarRoleType(m.GetValue()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemCalendarsItemAllowedCalendarSharingRolesWithUserResponse) SetValue(value []CalendarRoleType)() {
    m.value = value
}