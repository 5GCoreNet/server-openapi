/*
 * Nhss_EE
 *
 * HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_EE

// ReachabilityForDataConfiguration - Contains data needed for a Monitoring Configuration, specific to the 'Reachability for Data' event type
type ReachabilityForDataConfiguration struct {

	// indicating a time in seconds.
	MaximumLatency int32 `json:"maximumLatency,omitempty"`

	// indicating a time in seconds.
	MaximumResponseTime int32 `json:"maximumResponseTime,omitempty"`

	SuggestedPacketNumDl int32 `json:"suggestedPacketNumDl,omitempty"`
}
