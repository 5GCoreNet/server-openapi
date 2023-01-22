/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

type NrfInfoServedPcscfInfoListValueValue struct {

	AccessType []AccessType `json:"accessType,omitempty"`

	DnnList []string `json:"dnnList,omitempty"`

	// Fully Qualified Domain Name
	GmFqdn string `json:"gmFqdn,omitempty"`

	GmIpv4Addresses []string `json:"gmIpv4Addresses,omitempty"`

	GmIpv6Addresses []Ipv6Addr `json:"gmIpv6Addresses,omitempty"`

	// Fully Qualified Domain Name
	MwFqdn string `json:"mwFqdn,omitempty"`

	MwIpv4Addresses []string `json:"mwIpv4Addresses,omitempty"`

	MwIpv6Addresses []Ipv6Addr `json:"mwIpv6Addresses,omitempty"`

	ServedIpv4AddressRanges []Ipv4AddressRange `json:"servedIpv4AddressRanges,omitempty"`

	ServedIpv6PrefixRanges []Ipv6PrefixRange `json:"servedIpv6PrefixRanges,omitempty"`
}
