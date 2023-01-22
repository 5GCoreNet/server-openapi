/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_AMPolicyAuthorization

// AmEventsSubscRespData - Identifies the events the application subscribes to within an AM Policy Events Subscription subresource data. It may contain the notification of the already met events.
type AmEventsSubscRespData struct {

	// String providing an URI formatted according to RFC 3986.
	EventNotifUri string `json:"eventNotifUri"`

	Events []AmEventData `json:"events,omitempty"`

	// Contains the AM Policy Events Subscription resource identifier related to the event notification.
	AppAmContextId string `json:"appAmContextId,omitempty"`

	RepEvents []AmEventNotification `json:"repEvents"`
}
