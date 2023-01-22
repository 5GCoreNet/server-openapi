/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsmf_MBSSession

// ContextUpdateReqData - Data within ContextUpdate Request
type ContextUpdateReqData struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfcInstanceId string `json:"nfcInstanceId"`

	MbsSessionId MbsSessionId `json:"mbsSessionId"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId int32 `json:"areaSessionId,omitempty"`

	RequestedAction ContextUpdateAction `json:"requestedAction,omitempty"`

	// string with format 'bytes' as defined in OpenAPI
	DlTunnelInfo string `json:"dlTunnelInfo,omitempty"`

	N2MbsSmInfo N2MbsSmInfo `json:"n2MbsSmInfo,omitempty"`

	RanNodeId GlobalRanNodeId `json:"ranNodeId,omitempty"`

	LeaveInd bool `json:"leaveInd,omitempty"`
}
