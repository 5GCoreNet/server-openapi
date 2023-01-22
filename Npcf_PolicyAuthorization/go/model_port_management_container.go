/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

// PortManagementContainer - Contains the port management information container for a port.
type PortManagementContainer struct {

	// string with format 'bytes' as defined in OpenAPI
	PortManCont string `json:"portManCont"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNum int32 `json:"portNum"`
}
