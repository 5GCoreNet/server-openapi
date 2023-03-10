/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_TrafficInfluence

// AfResultInfo - Identifies the result of application layer handling.
type AfResultInfo struct {

	AfStatus AfResultStatus `json:"afStatus"`

	TrafficRoute *RouteToLocation `json:"trafficRoute,omitempty"`

	// If present and set to \"true\" it indicates that buffering of uplink traffic to the target DNAI is needed. 
	UpBuffInd bool `json:"upBuffInd,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
}
