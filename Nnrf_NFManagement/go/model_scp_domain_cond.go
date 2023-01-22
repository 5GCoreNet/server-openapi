/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// ScpDomainCond - Subscription to a set of NF or SCP or SEPP instances belonging to certain SCP domains 
type ScpDomainCond struct {

	ScpDomains []string `json:"scpDomains"`

	NfTypeList []NfType `json:"nfTypeList,omitempty"`
}
