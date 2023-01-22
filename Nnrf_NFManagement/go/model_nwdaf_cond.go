/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// NwdafCond - Subscription to a set of NF Instances (NWDAFs), identified by Analytics ID(s), S-NSSAI(s) or NWDAF Serving Area information, i.e. list of TAIs for which the NWDAF can provide analytics. 
type NwdafCond struct {

	ConditionType string `json:"conditionType"`

	AnalyticsIds []string `json:"analyticsIds,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	ServingNfTypeList []NfType `json:"servingNfTypeList,omitempty"`

	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`

	MlAnalyticsList []MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`
}
