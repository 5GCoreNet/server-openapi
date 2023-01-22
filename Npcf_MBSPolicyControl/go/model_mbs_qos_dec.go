/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_MBSPolicyControl

// MbsQosDec - Represents the parameters constituting an MBS QoS Decision.
type MbsQosDec struct {

	MbsQosId string `json:"mbsQosId"`

	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi,omitempty"`

	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MbrDl string `json:"mbrDl,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GbrDl string `json:"gbrDl,omitempty"`

	Arp Arp `json:"arp,omitempty"`

	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow int32 `json:"averWindow,omitempty"`

	// Represents MBS Maximum Data Burst Volume, expressed in Bytes.
	MbsMaxDataBurstVol int32 `json:"mbsMaxDataBurstVol,omitempty"`
}
