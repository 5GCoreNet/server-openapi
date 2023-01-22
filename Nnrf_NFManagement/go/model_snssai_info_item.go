/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// SnssaiInfoItem - Parameters supported by an NF for a given S-NSSAI Set of parameters supported by NF for a given S-NSSAI 
type SnssaiInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnInfoList []DnnInfoItem `json:"dnnInfoList"`
}
