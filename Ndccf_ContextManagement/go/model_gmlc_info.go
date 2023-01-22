/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// GmlcInfo - Information of a GMLC NF Instance
type GmlcInfo struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	GmlcNumbers []string `json:"gmlcNumbers,omitempty"`
}
