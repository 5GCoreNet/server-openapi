/*
 * Npcf_MBSPolicyAuthorization API
 *
 * MBS Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_MBSPolicyAuthorization

// MbsMediaCompRm - This data type is defined in the same way as the MbsMediaComp data type, but with the OpenAPI nullable property set to true. 
type MbsMediaCompRm struct {

	MbsMedCompNum int32 `json:"mbsMedCompNum"`

	MbsFlowDescs []string `json:"mbsFlowDescs,omitempty"`

	MbsSdfResPrio ReservPriority `json:"mbsSdfResPrio,omitempty"`

	MbsMediaInfo MbsMediaInfo `json:"mbsMediaInfo,omitempty"`

	QosRef string `json:"qosRef,omitempty"`

	MbsQoSReq MbsQoSReq `json:"mbsQoSReq,omitempty"`
}
