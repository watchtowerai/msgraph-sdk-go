package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse 
// Deprecated: This class is obsolete. Use getApplicablePolicyRequirementsPostResponse instead.
type EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse struct {
    EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponse
}
// NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse instantiates a new EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse and sets the default values.
func NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse()(*EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse) {
    m := &EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse{
        EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponse: *NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponse(),
    }
    return m
}
// CreateEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponse(), nil
}
// EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseable 
// Deprecated: This class is obsolete. Use getApplicablePolicyRequirementsPostResponse instead.
type EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsResponseable interface {
    EntitlementManagementAccessPackagesItemGetApplicablePolicyRequirementsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
