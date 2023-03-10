/*
 * Eees_ACREvents
 *
 * API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_ACREvents

import (
	"time"
)

// DiscoveredEas - Represents an EAS discovery information.
type DiscoveredEas struct {

	Eas EasProfile `json:"eas"`

	// string with format \"date-time\" as defined in OpenAPI.
	LifeTime time.Time `json:"lifeTime,omitempty"`
}
