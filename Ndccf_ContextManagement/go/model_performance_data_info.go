/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

import (
	"time"
)

// PerformanceDataInfo - Contains Performance Data Analytics related information collection.
type PerformanceDataInfo struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	UeIpAddr IpAddr `json:"ueIpAddr,omitempty"`

	IpTrafficFilter FlowInfo `json:"ipTrafficFilter,omitempty"`

	UserLoc UserLocation `json:"userLoc,omitempty"`

	AppLocs []string `json:"appLocs,omitempty"`

	AsAddr AddrFqdn `json:"asAddr,omitempty"`

	PerfData PerformanceData `json:"perfData"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}
