/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// UserLocation - At least one of eutraLocation, nrLocation and n3gaLocation shall be present. Several of them may be present. 
type UserLocation struct {

	EutraLocation EutraLocation `json:"eutraLocation,omitempty"`

	NrLocation NrLocation `json:"nrLocation,omitempty"`

	N3gaLocation N3gaLocation `json:"n3gaLocation,omitempty"`

	UtraLocation UtraLocation `json:"utraLocation,omitempty"`

	GeraLocation GeraLocation `json:"geraLocation,omitempty"`
}
