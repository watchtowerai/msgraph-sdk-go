package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse struct {
    EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse
}
// NewEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse instantiates a new EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse()(*EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse) {
    m := &EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse{
        EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse: *NewEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse(), nil
}
// EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable interface {
    EntitlementManagementRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
