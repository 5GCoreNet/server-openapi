/*
 * CAPIF_API_Provider_Management_API
 *
 * API for API provider domain functions management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_API_Provider_Management_API

// RegistrationInformation - Represents registration information of an individual API provider domain function. 
type RegistrationInformation struct {

	// Public Key of API Provider domain function.
	ApiProvPubKey string `json:"apiProvPubKey"`

	// API provider domain function's client certificate
	ApiProvCert string `json:"apiProvCert,omitempty"`
}
