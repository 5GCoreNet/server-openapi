/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsSDM

import (
	"time"
)

// ImsSdmSubscription - A subscription to notifications of data change
type ImsSdmSubscription struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`

	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`

	MonitoredResourceUris []string `json:"monitoredResourceUris"`

	// string with format 'date-time' as defined in OpenAPI.
	Expires time.Time `json:"expires,omitempty"`
}
