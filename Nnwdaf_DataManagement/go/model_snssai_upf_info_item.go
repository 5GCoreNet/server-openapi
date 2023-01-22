/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// SnssaiUpfInfoItem - Set of parameters supported by UPF for a given S-NSSAI
type SnssaiUpfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList"`

	RedundantTransport bool `json:"redundantTransport,omitempty"`
}
