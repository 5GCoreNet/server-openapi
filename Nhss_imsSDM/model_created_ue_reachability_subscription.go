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

// CreatedUeReachabilitySubscription - Contains the response data returned by HSS after the subscription to  notifications of UE reachability for IP was created 
type CreatedUeReachabilitySubscription struct {

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry"`
}
