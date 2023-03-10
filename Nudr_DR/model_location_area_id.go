/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// LocationAreaId - Contains a Location area identification as defined in 3GPP TS 23.003, clause 4.1.
type LocationAreaId struct {

	PlmnId PlmnId1 `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac"`
}
