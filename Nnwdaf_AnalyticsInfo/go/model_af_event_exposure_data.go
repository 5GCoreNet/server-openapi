/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// AfEventExposureData - AF Event Exposure data managed by a given NEF Instance
type AfEventExposureData struct {

	AfEvents []AfEvent `json:"afEvents"`

	AfIds []string `json:"afIds,omitempty"`

	AppIds []string `json:"appIds,omitempty"`
}
