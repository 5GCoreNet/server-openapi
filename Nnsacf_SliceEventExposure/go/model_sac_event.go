/*
 * Nnsacf_SliceEventExposure
 *
 * Nnsacf_SliceEventExposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnsacf_SliceEventExposure

// SacEvent - Describes an event to be subscribed
type SacEvent struct {

	EventType SacEventType `json:"eventType"`

	EventTrigger SacEventTrigger `json:"eventTrigger,omitempty"`

	EventFilter []Snssai `json:"eventFilter"`

	// indicating a time in seconds.
	NotificationPeriod int32 `json:"notificationPeriod,omitempty"`

	NotifThreshold SacInfo `json:"notifThreshold,omitempty"`

	ImmediateFlag bool `json:"immediateFlag,omitempty"`
}
