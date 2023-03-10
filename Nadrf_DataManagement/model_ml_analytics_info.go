/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// MlAnalyticsInfo - ML Analytics Filter information supported by the Nnwdaf_MLModelProvision service
type MlAnalyticsInfo struct {

	MlAnalyticsIds []NwdafEvent `json:"mlAnalyticsIds,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`
}
