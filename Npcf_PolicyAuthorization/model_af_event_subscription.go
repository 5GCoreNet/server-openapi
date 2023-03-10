/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// AfEventSubscription - Describes the event information delivered in the subscription.
type AfEventSubscription struct {

	Event AfEvent `json:"event"`

	NotifMethod AfNotifMethod `json:"notifMethod,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	// indicating a time in seconds.
	WaitTime int32 `json:"waitTime,omitempty"`
}
