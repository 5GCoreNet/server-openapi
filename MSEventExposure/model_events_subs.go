/*
 * 3gpp-ms-event-exposure
 *
 * API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSEventExposure

// EventsSubs - Represents an event to be subscribed and the related event filter information.
type EventsSubs struct {

	Event AfEvent `json:"event"`

	EventFilter EventFilter `json:"eventFilter"`
}
