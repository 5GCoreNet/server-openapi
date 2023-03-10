/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

// SacEventStatus - Contains the network slice status information in terms of the current number of UEs registered  with a network slice, the current number of PDU Sessions established on a network slice or both. 
type SacEventStatus struct {

	ReachedNumUes SacInfo `json:"reachedNumUes,omitempty"`

	ReachedNumPduSess SacInfo `json:"reachedNumPduSess,omitempty"`
}
