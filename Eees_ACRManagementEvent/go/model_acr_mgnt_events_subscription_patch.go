/*
 * EES ACR Management Event_API
 *
 * API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_ACRManagementEvent

// AcrMgntEventsSubscriptionPatch - Represents a modification request of Individual ACR Management Events Subscription. 
type AcrMgntEventsSubscriptionPatch struct {

	// The subscribed ACR management events.
	EventSubscs []AcrMgntEventSubsc `json:"eventSubscs,omitempty"`

	EvtReq ReportingInformation `json:"evtReq,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}