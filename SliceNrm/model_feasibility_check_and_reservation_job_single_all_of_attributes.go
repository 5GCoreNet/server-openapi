/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type FeasibilityCheckAndReservationJobSingleAllOfAttributes struct {

	Profile *OneOfSliceProfileServiceProfile `json:"profile,omitempty"`

	// An attribute represents MnS consumer's requirements for resource reservation.
	ResourceReservation bool `json:"resourceReservation,omitempty"`

	// An attribute represents MnS consumer's request for recommended network slice related requirements.
	RecommendationRequest bool `json:"recommendationRequest,omitempty"`

	// An attribute which specifes MnS consuner's requirements for the validity period of the resource reservation.
	RequestedReservationExpiration string `json:"requestedReservationExpiration,omitempty"`

	ProcessMonitor ProcessMonitor `json:"processMonitor,omitempty"`

	FeasibilityResult FeasibilityResult `json:"feasibilityResult,omitempty"`

	// An attribute that specifies the additional reason information if the feasibility check result is infeasible.The detailed ENUM value is FFS. 
	InFeasibleReason string `json:"inFeasibleReason,omitempty"`

	ResourceReservationStatus ResourceReservationStatus `json:"resourceReservationStatus,omitempty"`

	// An attribute that specifies the additional reason information if the reservation is failed. 
	ReservationFailureReason string `json:"reservationFailureReason,omitempty"`

	// An attribute which specifes the actual validity period of the resource reservation.
	ReservationExpiration string `json:"reservationExpiration,omitempty"`

	// An attribute that specifies the recommended network slicing related requirements (i.e. ServiceProfile and SliceProfile information) which can be supported by the MnS producer.
	RecommendedRequirements string `json:"recommendedRequirements,omitempty"`
}
