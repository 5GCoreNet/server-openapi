/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// NfGroupListCond - Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {

	ConditionType string `json:"conditionType"`

	NfType string `json:"nfType"`

	NfGroupIdList []string `json:"nfGroupIdList"`
}
