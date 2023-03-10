/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnef_SMContext

// SmContextUpdateData - Representation of the updates to apply to the Individual SM context.
type SmContextUpdateData struct {

	// String providing an URI formatted according to RFC 3986.
	DlNiddEndPoint string `json:"dlNiddEndPoint,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri,omitempty"`

	SmContextConfig SmContextConfiguration `json:"smContextConfig,omitempty"`
}
