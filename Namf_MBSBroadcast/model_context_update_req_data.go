/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_MBSBroadcast

// ContextUpdateReqData - Data within ContextUpdate Request
type ContextUpdateReqData struct {

	MbsServiceArea MbsServiceArea `json:"mbsServiceArea,omitempty"`

	MbsServiceAreaInfoList []MbsServiceAreaInfo `json:"mbsServiceAreaInfoList,omitempty"`

	N2MbsSmInfo N2MbsSmInfo `json:"n2MbsSmInfo,omitempty"`

	RanIdList []GlobalRanNodeId `json:"ranIdList,omitempty"`

	NoNgapSignallingInd bool `json:"noNgapSignallingInd,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri,omitempty"`

	// indicating a time in seconds.
	MaxResponseTime int32 `json:"maxResponseTime,omitempty"`

	N2MbsInfoChangeInd bool `json:"n2MbsInfoChangeInd,omitempty"`
}
