package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse struct {
    AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnGetResponse
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse()(*AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse{
        AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnGetResponse: *NewAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponse(), nil
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnResponseable interface {
    AccessReviewsDefinitionsItemInstancesItemDecisionsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
