/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// EcsServerAddr - Contains the Edge Configuration Server Address Configuration Information as defined in clause 5.2.3.6.1 of 3GPP TS 23.502. 
type EcsServerAddr struct {

	EcsFqdnList []string `json:"ecsFqdnList,omitempty"`

	EcsIpAddressList []IpAddr `json:"ecsIpAddressList,omitempty"`

	EcsUriList []string `json:"ecsUriList,omitempty"`

	EcsProviderId string `json:"ecsProviderId,omitempty"`
}
