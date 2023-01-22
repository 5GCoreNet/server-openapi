/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EventExposure

// NefEventSubs - Represents an event to be subscribed and the related event filter information.
type NefEventSubs struct {

	Event NefEvent `json:"event"`

	EventFilter NefEventFilter `json:"eventFilter,omitempty"`
}
