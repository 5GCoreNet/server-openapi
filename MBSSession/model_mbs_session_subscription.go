/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSSession

import (
	"time"
)

// MbsSessionSubscription - MBS session subscription
type MbsSessionSubscription struct {

	MbsSessionId MbsSessionId `json:"mbsSessionId,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId int32 `json:"areaSessionId,omitempty"`

	EventList []MbsSessionEvent `json:"eventList"`

	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri"`

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ExpiryTime time.Time `json:"expiryTime,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfcInstanceId string `json:"nfcInstanceId,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	MbsSessionSubscUri string `json:"mbsSessionSubscUri,omitempty"`
}
