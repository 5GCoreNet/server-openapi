/*
 * CAPIF_Publish_Service_API
 *
 * API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CAPIF_Publish_Service_API

// ShareableInformation - Indicates whether the service API and/or the service API category can be shared to the list of CAPIF provider domains. 
type ShareableInformation struct {

	// Set to \"true\" indicates that the service API and/or the service API category can be shared to the list of CAPIF provider domain information. Otherwise set to \"false\". 
	IsShareable bool `json:"isShareable"`

	// List of CAPIF provider domains to which the service API information to be shared. 
	CapifProvDoms []string `json:"capifProvDoms,omitempty"`
}
