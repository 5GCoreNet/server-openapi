/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// MulticastTransportAddressChangeInfo - Multicast Transport Address Change Information
type MulticastTransportAddressChangeInfo struct {

	LlSsm Ssm `json:"llSsm"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	CTeid int32 `json:"cTeid"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId int32 `json:"areaSessionId,omitempty"`
}
