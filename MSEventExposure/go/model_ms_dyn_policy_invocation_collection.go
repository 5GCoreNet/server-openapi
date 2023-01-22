/*
 * 3gpp-ms-event-exposure
 *
 * API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MSEventExposure

// MsDynPolicyInvocationCollection - Contains the Media Streaming Dynamic Policy invocation collected for an UE Application via AF. 
type MsDynPolicyInvocationCollection struct {

	MsDynPlyInvocs []DynamicPolicy `json:"msDynPlyInvocs"`
}
