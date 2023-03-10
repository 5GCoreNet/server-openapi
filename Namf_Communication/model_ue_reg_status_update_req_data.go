/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// UeRegStatusUpdateReqData - Data within a UE registration status update request to indicate a completion of transferring at a target AMF
type UeRegStatusUpdateReqData struct {

	TransferStatus UeContextTransferStatus `json:"transferStatus"`

	ToReleaseSessionList []int32 `json:"toReleaseSessionList,omitempty"`

	PcfReselectedInd bool `json:"pcfReselectedInd,omitempty"`

	SmfChangeInfoList []SmfChangeInfo `json:"smfChangeInfoList,omitempty"`

	AnalyticsNotUsedList []string `json:"analyticsNotUsedList,omitempty"`

	ToReleaseSessionInfo []ReleaseSessionInfo `json:"toReleaseSessionInfo,omitempty"`
}
