/*
 * 3gpp-ms-event-exposure
 *
 * API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSEventExposure

// ServiceExperienceInfoPerApp - Contains service experience information associated with an application.
type ServiceExperienceInfoPerApp struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	AppServerIns AddrFqdn `json:"appServerIns,omitempty"`

	SvcExpPerFlows []ServiceExperienceInfoPerFlow `json:"svcExpPerFlows"`

	Gpsis []string `json:"gpsis,omitempty"`

	Supis []string `json:"supis,omitempty"`
}
