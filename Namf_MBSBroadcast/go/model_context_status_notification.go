/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_MBSBroadcast

// ContextStatusNotification - Data within ContextStatusNotify Request
type ContextStatusNotification struct {

	MbsSessionId MbsSessionId `json:"mbsSessionId"`

	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId int32 `json:"areaSessionId,omitempty"`

	N2MbsSmInfoList []N2MbsSmInfo `json:"n2MbsSmInfoList,omitempty"`

	OperationEvents []OperationEvent `json:"operationEvents,omitempty"`

	OperationStatus OperationStatus `json:"operationStatus,omitempty"`
}
