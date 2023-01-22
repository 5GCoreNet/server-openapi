/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// N2InterfaceAmfInfo - AMF N2 interface information
type N2InterfaceAmfInfo struct {

	Ipv4EndpointAddress []string `json:"ipv4EndpointAddress,omitempty"`

	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`

	// Fully Qualified Domain Name
	AmfName string `json:"amfName,omitempty"`
}
