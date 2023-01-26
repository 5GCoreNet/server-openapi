/*
 * Nudm_SSAU
 *
 * Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SSAU

import (
	"time"
)

// ServiceSpecificAuthorizationData - Authorization Response for a specific service.
type ServiceSpecificAuthorizationData struct {

	AuthorizationUeId AuthorizationUeId `json:"authorizationUeId,omitempty"`

	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGroupId string `json:"extGroupId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime,omitempty"`
}