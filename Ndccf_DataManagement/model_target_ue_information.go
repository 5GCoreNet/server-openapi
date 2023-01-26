/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// TargetUeInformation - Identifies the target UE information.
type TargetUeInformation struct {

	AnyUe bool `json:"anyUe,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	IntGroupIds []string `json:"intGroupIds,omitempty"`
}