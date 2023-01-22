/*
 * TS 29.122 Common Data Types
 *
 * Data types applicable to several APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// EthFlowInfo - Represents Ethernet flow information.
type EthFlowInfo struct {

	// Indicates the Ethernet flow identifier.
	FlowId int32 `json:"flowId"`

	// Indicates the packet filters of the Ethernet flow. It shall contain UL and/or DL Ethernet flow description. 
	EthFlowDescriptions []EthFlowDescription `json:"ethFlowDescriptions,omitempty"`
}
