/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// NfGroupListCond - Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {

	ConditionType string `json:"conditionType"`

	NfType string `json:"nfType"`

	NfGroupIdList []string `json:"nfGroupIdList"`
}