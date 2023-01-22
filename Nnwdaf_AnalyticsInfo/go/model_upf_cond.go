/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// UpfCond - Subscription to a set of NF Instances (UPFs), able to serve a certain service area (i.e. SMF serving area or TAI list) 
type UpfCond struct {

	ConditionType string `json:"conditionType"`

	SmfServingArea []string `json:"smfServingArea,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`
}