/*
 * CAPIF_API_Provider_Management_API
 *
 * API for API provider domain functions management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_API_Provider_Management_API

// ApiProviderFunctionDetails - Represents API provider domain function's details.
type ApiProviderFunctionDetails struct {

	// API provider domain functionID assigned by the CAPIF core function to the API provider domain function while registering/updating the API provider domain. Shall not be present in the HTTP POST request from the API management function to the CAPIF core function, to register itself. Shall be present in all other HTTP requests and responses. 
	ApiProvFuncId string `json:"apiProvFuncId,omitempty"`

	RegInfo RegistrationInformation `json:"regInfo"`

	ApiProvFuncRole ApiProviderFuncRole `json:"apiProvFuncRole"`

	// Generic information related to the API provider domain function such as details of the API provider applications. 
	ApiProvFuncInfo string `json:"apiProvFuncInfo,omitempty"`
}
