/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFDiscovery

// DnnSmfInfoItem - Set of parameters supported by SMF for a given DNN
type DnnSmfInfoItem struct {

	Dnn DnnSmfInfoItemDnn `json:"dnn"`

	DnaiList []DnnSmfInfoItemDnaiListInner `json:"dnaiList,omitempty"`
}
