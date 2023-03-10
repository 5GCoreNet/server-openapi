/*
 * Npcf_MBSPolicyAuthorization API
 *
 * MBS Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_MBSPolicyAuthorization

// Ssm - Source specific IP multicast address
type Ssm struct {

	SourceIpAddr IpAddr `json:"sourceIpAddr"`

	DestIpAddr IpAddr `json:"destIpAddr"`
}
