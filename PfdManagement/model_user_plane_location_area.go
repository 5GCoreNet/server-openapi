/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_PfdManagement

// UserPlaneLocationArea - Represents location area(s) of the user plane functions which are unable to enforce the provisioned PFD(s) successfully.
type UserPlaneLocationArea struct {

	LocationArea LocationArea `json:"locationArea,omitempty"`

	LocationArea5G LocationArea5G `json:"locationArea5G,omitempty"`

	// Identifies a list of DNAI which the user plane functions support.
	Dnais []string `json:"dnais,omitempty"`
}
