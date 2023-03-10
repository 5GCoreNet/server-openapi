/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// HalTemplate - Hypertext Application Language (HAL) template contains the extended 3GPP hypermedia format. 
type HalTemplate struct {

	// A human-readable string that can be used to identify this template
	Title string `json:"title,omitempty"`

	Method HttpMethod `json:"method"`

	// The media type that should be used for the corresponding request. If the attribute is missing, or contains an unrecognized value, the client should act as if the  contentType is set to \"application/json\". 
	ContentType string `json:"contentType,omitempty"`

	// The properties that should be included in the body of the corresponding request.  If the contentType attribute is set to \"application/json\", then this attribute  describes the attributes of the JSON object of the body. 
	Properties []Property `json:"properties,omitempty"`
}
