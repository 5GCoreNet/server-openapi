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

// NiddConfigurationPatch - Represents the parameters to update a NIDD configuration.
type NiddConfigurationPatch struct {

	// string with format \"date-time\" as defined in OpenAPI with \"nullable=true\" property.
	Duration *time.Time `json:"duration,omitempty"`

	// Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false). 
	ReliableDataService *bool `json:"reliableDataService,omitempty"`

	// Indicates the static port configuration that is used for reliable data transfer between specific applications using RDS (as defined in clause 5.2.4 and 5.2.5 of 3GPP TS 24.250).
	RdsPorts []RdsPort `json:"rdsPorts,omitempty"`

	PdnEstablishmentOption PdnEstablishmentOptionsRm `json:"pdnEstablishmentOption,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}
