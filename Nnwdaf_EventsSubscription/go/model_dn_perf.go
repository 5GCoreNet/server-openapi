/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

// DnPerf - Represents DN performance for the application.
type DnPerf struct {

	AppServerInsAddr AddrFqdn `json:"appServerInsAddr,omitempty"`

	UpfInfo UpfInformation `json:"upfInfo,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai,omitempty"`

	PerfData PerfData `json:"perfData"`

	SpatialValidCon NetworkAreaInfo `json:"spatialValidCon,omitempty"`

	TemporalValidCon TimeWindow `json:"temporalValidCon,omitempty"`
}
