/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsf_MBSUserDataIngestSession

// Ssm - Source specific IP multicast address
type Ssm struct {

	SourceIpAddr IpAddr `json:"sourceIpAddr"`

	DestIpAddr IpAddr `json:"destIpAddr"`
}
