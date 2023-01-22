/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// MbUpfInfo - Information of an MB-UPF NF Instance
type MbUpfInfo struct {

	SNssaiMbUpfInfoList []SnssaiUpfInfoItem `json:"sNssaiMbUpfInfoList"`

	MbSmfServingArea []string `json:"mbSmfServingArea,omitempty"`

	InterfaceMbUpfInfoList []InterfaceUpfInfoItem `json:"interfaceMbUpfInfoList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	SupportedPfcpFeatures string `json:"supportedPfcpFeatures,omitempty"`
}
