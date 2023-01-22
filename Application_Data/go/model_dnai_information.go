/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Application_Data

// DnaiInformation - Represents DNAI information.
type DnaiInformation struct {

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`

	DnsServIds []DnsServerIdentifier `json:"dnsServIds,omitempty"`

	EasIpAddrs []IpAddr `json:"easIpAddrs,omitempty"`
}
