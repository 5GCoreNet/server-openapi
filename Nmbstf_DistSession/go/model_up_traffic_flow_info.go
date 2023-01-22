/*
 * Nmbstf-distsession
 *
 * MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbstf_DistSession

// UpTrafficFlowInfo - IP Multicast Address and Port Number
type UpTrafficFlowInfo struct {

	DestIpAddr IpAddr `json:"destIpAddr"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber,omitempty"`
}
