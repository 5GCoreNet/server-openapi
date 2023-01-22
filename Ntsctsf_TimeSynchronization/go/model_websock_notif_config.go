/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ntsctsf_TimeSynchronization

// WebsockNotifConfig - Represents the configuration information for the delivery of notifications over Websockets.
type WebsockNotifConfig struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	WebsocketUri string `json:"websocketUri,omitempty"`

	// Set by the SCS/AS to indicate that the Websocket delivery is requested.
	RequestWebsocketUri bool `json:"requestWebsocketUri,omitempty"`
}
