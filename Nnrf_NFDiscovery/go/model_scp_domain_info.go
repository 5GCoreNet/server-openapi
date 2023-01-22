/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFDiscovery

// ScpDomainInfo - SCP Domain specific information
type ScpDomainInfo struct {

	// Fully Qualified Domain Name
	ScpFqdn string `json:"scpFqdn,omitempty"`

	ScpIpEndPoints []IpEndPoint `json:"scpIpEndPoints,omitempty"`

	ScpPrefix string `json:"scpPrefix,omitempty"`

	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\". 
	ScpPorts map[string]int32 `json:"scpPorts,omitempty"`
}
