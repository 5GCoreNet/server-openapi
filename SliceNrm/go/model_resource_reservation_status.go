/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm
// ResourceReservationStatus : An attribute which specifies the resource reservation result for the feasibility check job.
type ResourceReservationStatus string

// List of ResourceReservationStatus
const (
	RESERVED ResourceReservationStatus = "RESERVED"
	UNRESERVED ResourceReservationStatus = "UNRESERVED"
	USED ResourceReservationStatus = "USED"
)
