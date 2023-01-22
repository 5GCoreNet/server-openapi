/*
 * 3gpp-chargeable-party
 *
 * API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ChargeableParty

// EthFlowDescription - Identifies an Ethernet flow.
type EthFlowDescription struct {

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddr string `json:"destMacAddr,omitempty"`

	EthType string `json:"ethType"`

	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`

	FDir FlowDirection `json:"fDir,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SourceMacAddr string `json:"sourceMacAddr,omitempty"`

	VlanTags []string `json:"vlanTags,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty"`
}
