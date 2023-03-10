/*
 * CAPIF_API_Provider_Management_API
 *
 * API for API provider domain functions management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_API_Provider_Management_API

// ApiProviderEnrolmentDetails - Represents an API provider domain's enrolment details.
type ApiProviderEnrolmentDetails struct {

	// API provider domain ID assigned by the CAPIF core function to the API management function while registering the API provider domain. Shall not be present in the HTTP POST request from the API Management function to the CAPIF core function, to on-board itself. Shall be present in all other HTTP requests and responses. 
	ApiProvDomId string `json:"apiProvDomId,omitempty"`

	// Security information necessary for the CAPIF core function to validate the registration of the API provider domain. Shall be present in HTTP POST request from API management function to CAPIF core function for API provider domain registration. 
	RegSec string `json:"regSec"`

	// A list of individual API provider domain functions details. When included by the API management function in the HTTP request message, it lists the API provider domain functions that the API management function intends to register/update in registration or update registration procedure. When included by the CAPIF core function in the HTTP response message, it lists the API domain functions details that are registered or updated successfully. 
	ApiProvFuncs []ApiProviderFunctionDetails `json:"apiProvFuncs,omitempty"`

	// Generic information related to the API provider domain such as details of the API provider applications. 
	ApiProvDomInfo string `json:"apiProvDomInfo,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`

	// Registration or update specific failure information of failed API provider domain function registrations.Shall be present in the HTTP response body if atleast one of the API provider domain function registration or update registration fails. 
	FailReason string `json:"failReason,omitempty"`
}
