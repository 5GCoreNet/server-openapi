/*
 * 3gpp-eas-deployment
 *
 * API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EASDeployment

// DnsServerIdentifier - Represents DNS server identifier (consisting of IP address and port).
type DnsServerIdentifier struct {

	DnsServIpAddr IpAddr `json:"dnsServIpAddr"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}
