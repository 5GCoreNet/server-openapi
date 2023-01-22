/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsmf_MBSSession

// Ssm - Source specific IP multicast address
type Ssm struct {

	SourceIpAddr IpAddr `json:"sourceIpAddr"`

	DestIpAddr IpAddr `json:"destIpAddr"`
}
