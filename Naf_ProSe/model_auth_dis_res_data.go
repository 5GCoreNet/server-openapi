/*
 * Naf_ProSe API
 *
 * Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Naf_ProSe

// AuthDisResData - Represents the obtained authorization Data for a UE of a 5G ProSe Direct Discovery  request. 
type AuthDisResData struct {

	AuthResponseType AuthResponseType `json:"authResponseType"`

	ProseAppCodeSuffixPool ProseApplicationCodeSuffixPool `json:"proseAppCodeSuffixPool,omitempty"`

	Pduids []string `json:"pduids,omitempty"`

	RestrictedCodeSuffixPool []RestrictedCodeSuffixPool `json:"restrictedCodeSuffixPool,omitempty"`

	ProseAppMasks []string `json:"proseAppMasks,omitempty"`

	ProSeRestrictedMasks []string `json:"proSeRestrictedMasks,omitempty"`

	// Contains the Application Level Container.
	ResAppLevelContainer string `json:"resAppLevelContainer,omitempty"`

	TargetDataSet []TargetData `json:"targetDataSet,omitempty"`

	// Contains the PDUID.
	TargetPduid string `json:"targetPduid,omitempty"`

	// Contains the metadata.
	MetaData string `json:"metaData,omitempty"`
}
