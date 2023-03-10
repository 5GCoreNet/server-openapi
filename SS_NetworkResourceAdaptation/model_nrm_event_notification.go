/*
 * SS_NetworkResourceAdaptation
 *
 * SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_NetworkResourceAdaptation

import (
	"time"
)

// NrmEventNotification - Represents a notification on an individual User Plane event.
type NrmEventNotification struct {

	Event NrmEvent `json:"event"`

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`

	DeliveryMode DeliveryMode `json:"deliveryMode,omitempty"`

	StreamIds []string `json:"streamIds,omitempty"`
}
