/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// SnssaiTsctsfInfoItem - Set of parameters supported by TSCTSF for a given S-NSSAI
type SnssaiTsctsfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnInfoList []DnnTsctsfInfoItem `json:"dnnInfoList"`
}
