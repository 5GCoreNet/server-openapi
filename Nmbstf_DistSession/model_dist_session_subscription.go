/*
 * Nmbstf-distsession
 *
 * MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbstf_DistSession

import (
	"time"
)

// DistSessionSubscription - Data within the Status Subscription
type DistSessionSubscription struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfcInstanceId string `json:"nfcInstanceId,omitempty"`

	EventList []DistSessionEventType `json:"eventList"`

	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri"`

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ExpiryTime time.Time `json:"expiryTime,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	DistSessionSubscUri string `json:"distSessionSubscUri,omitempty"`
}
