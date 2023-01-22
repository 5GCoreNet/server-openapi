/*
 * 3gpp-ms-event-exposure
 *
 * API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MSEventExposure

// MsNetAssInvocationCollection - Contains the Media Streaming Network Assistance invocation collected for an UE Application  via AF. 
type MsNetAssInvocationCollection struct {

	MsNetAssInvocs []NetworkAssistanceSession `json:"msNetAssInvocs"`
}
