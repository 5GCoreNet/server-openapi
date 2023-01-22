/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

type NrfInfoServedBsfInfoValue struct {

	DnnList []string `json:"dnnList,omitempty"`

	IpDomainList []string `json:"ipDomainList,omitempty"`

	Ipv4AddressRanges []Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`

	Ipv6PrefixRanges []Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`

	// Fully Qualified Domain Name
	RxDiamHost string `json:"rxDiamHost,omitempty"`

	// Fully Qualified Domain Name
	RxDiamRealm string `json:"rxDiamRealm,omitempty"`

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`
}
