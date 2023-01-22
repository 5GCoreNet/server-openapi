/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASDiscovery

import (
	"time"
)

// EasDiscoverySubscriptionPatch - Represents an Individual EAS Discovery Subscription resource.
type EasDiscoverySubscriptionPatch struct {

	EasDiscoveryFilter EasDiscoveryFilter `json:"easDiscoveryFilter,omitempty"`

	EasDynInfoFilter EasDynamicInfoFilter `json:"easDynInfoFilter,omitempty"`

	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EasSvcContinuity []AcrScenario `json:"easSvcContinuity,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime time.Time `json:"expTime,omitempty"`
}
