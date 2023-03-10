/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NIDD

import (
	"time"
)

// NiddDownlinkDataDeliveryFailure - Represents information related to a failure delivery result.
type NiddDownlinkDataDeliveryFailure struct {

	ProblemDetail ProblemDetails `json:"problemDetail"`

	// string with format \"date-time\" as defined in OpenAPI.
	RequestedRetransmissionTime time.Time `json:"requestedRetransmissionTime,omitempty"`
}
