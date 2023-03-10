/*
 * CAPIF_Discover_Service_API
 *
 * API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Discover_Service_API

// DiscoveredApis - Represents a list of APIs currently registered in the CAPIF core function and satisfying a number of filter criteria provided by the API consumer. 
type DiscoveredApis struct {

	// Description of the service API as published by the service. Each service API description shall include AEF profiles matching the filter criteria. 
	ServiceAPIDescriptions []ServiceApiDescription `json:"serviceAPIDescriptions,omitempty"`
}
