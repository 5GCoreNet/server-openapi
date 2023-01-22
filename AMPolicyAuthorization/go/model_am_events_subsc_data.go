/*
 * 3gpp-am-policyauthorization
 *
 * API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AMPolicyAuthorization

// AmEventsSubscData - It represents the AM Policy Events Subscription subresource and identifies the events the application subscribes to.
type AmEventsSubscData struct {

	// String providing an URI formatted according to RFC 3986.
	EventNotifUri string `json:"eventNotifUri"`

	Events []AmEventData `json:"events,omitempty"`
}
