/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_Events

// SealEventSubscriptionPatch - Represents the partial update of individual SEAL Event Subscription resource.
type SealEventSubscriptionPatch struct {

	// Subscribed events.
	EventSubs []EventSubscription `json:"eventSubs,omitempty"`

	EventReq ReportingInformation `json:"eventReq,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}