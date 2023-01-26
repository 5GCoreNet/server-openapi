/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsSDM

// PriorityLevels - Namespaces and priority levels allowed for the IMS public Identity
type PriorityLevels struct {

	ServicePriorityLevelList []string `json:"servicePriorityLevelList"`

	ServicePriorityLevel int32 `json:"servicePriorityLevel,omitempty"`
}