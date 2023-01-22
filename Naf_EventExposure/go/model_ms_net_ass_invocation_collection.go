/*
 * Naf_EventExposure
 *
 * AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Naf_EventExposure

// MsNetAssInvocationCollection - Contains the Media Streaming Network Assistance invocation collected for an UE Application  via AF. 
type MsNetAssInvocationCollection struct {

	MsNetAssInvocs []NetworkAssistanceSession `json:"msNetAssInvocs"`
}
