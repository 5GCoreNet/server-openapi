/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

// NwdafSubscription - Individual NWDAF subscription identified by the subscription Id.
type NwdafSubscription struct {

	// String providing an URI formatted according to RFC 3986.
	NwdafEvtSubsServiceUri string `json:"nwdafEvtSubsServiceUri"`

	NwdafEventsSubscription NnwdafEventsSubscription `json:"nwdafEventsSubscription"`
}
