/*
 * Eees_ACREvents
 *
 * API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_ACREvents

import (
	"time"
)

// EdnConfigInfo - Represents the EDN information.
type EdnConfigInfo struct {

	EdnConInfo EdnConInfo `json:"ednConInfo"`

	// Contains the list of EESs of the EDN.
	Eess []EesInfo `json:"eess"`

	// string with format \"date-time\" as defined in OpenAPI.
	LifeTime time.Time `json:"lifeTime,omitempty"`
}
