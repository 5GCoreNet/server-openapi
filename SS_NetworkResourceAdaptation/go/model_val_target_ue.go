/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_NetworkResourceAdaptation

// ValTargetUe - Represents information identifying a VAL user ID or a VAL UE ID.
type ValTargetUe struct {

	// Unique identifier of a VAL user.
	ValUserId string `json:"valUserId,omitempty"`

	// Unique identifier of a VAL UE.
	ValUeId string `json:"valUeId,omitempty"`
}
