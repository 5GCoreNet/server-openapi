/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// GmlcInfo - Information of a GMLC NF Instance
type GmlcInfo struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	GmlcNumbers []string `json:"gmlcNumbers,omitempty"`
}
