/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// WlanPerSsIdPerformanceInfo - The WLAN performance per SSID.
type WlanPerSsIdPerformanceInfo struct {

	SsId string `json:"ssId"`

	WlanPerTsInfos []WlanPerTsPerformanceInfo `json:"wlanPerTsInfos"`
}
