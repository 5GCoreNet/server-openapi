/*
 * TS 29.122 Common Data Types
 *
 * Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// FlowInfo - Represents IP flow information.
type FlowInfo struct {

	// Indicates the IP flow identifier.
	FlowId int32 `json:"flowId"`

	// Indicates the packet filters of the IP flow. Refer to clause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description. 
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
}
