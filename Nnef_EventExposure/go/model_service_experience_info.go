/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EventExposure

// ServiceExperienceInfo - Contains service experience information associated with an application.
type ServiceExperienceInfo struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	Supis []string `json:"supis,omitempty"`

	SvcExpPerFlows []ServiceExperienceInfoPerFlow `json:"svcExpPerFlows"`
}