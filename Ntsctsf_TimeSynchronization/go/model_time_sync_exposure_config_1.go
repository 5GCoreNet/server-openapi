/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ntsctsf_TimeSynchronization

// TimeSyncExposureConfig1 - Contains the Time Synchronization Configuration parameters.
type TimeSyncExposureConfig1 struct {

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UpNodeId int32 `json:"upNodeId"`

	ReqPtpIns PtpInstance `json:"reqPtpIns"`

	// Indicates that the AF requests 5GS to act as a grandmaster for PTP or gPTP if it is included and set to true. 
	GmEnable bool `json:"gmEnable,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	GmPrio int32 `json:"gmPrio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeDom int32 `json:"timeDom"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeSyncErrBdgt int32 `json:"timeSyncErrBdgt,omitempty"`

	// Notification Correlation ID assigned by the NF service consumer.
	ConfigNotifId string `json:"configNotifId"`

	// String providing an URI formatted according to RFC 3986.
	ConfigNotifUri string `json:"configNotifUri"`

	TempValidity TemporalValidity `json:"tempValidity,omitempty"`
}
