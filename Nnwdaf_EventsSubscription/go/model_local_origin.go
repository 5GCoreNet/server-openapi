/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

// LocalOrigin - Indicates a Local origin in a reference system
type LocalOrigin struct {

	CoordinateId string `json:"coordinateId,omitempty"`

	Point GeographicalCoordinates `json:"point,omitempty"`
}
