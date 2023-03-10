/*
 * Namf_Location
 *
 * AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Location

// UePrivacyRequirements - UE privacy requirements from (H)GMLC to the serving AMF or VGMLC(in the roaming case) for the target UE
type UePrivacyRequirements struct {

	LcsServiceAuthInfo LcsServiceAuth `json:"lcsServiceAuthInfo,omitempty"`

	CodeWordCheck bool `json:"codeWordCheck,omitempty"`
}
