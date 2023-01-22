/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EventExposure

import (
	"time"
)

// UeTrajectoryInfo - Contains UE trajectory information.
type UeTrajectoryInfo struct {

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`

	Location UserLocation `json:"location"`
}
