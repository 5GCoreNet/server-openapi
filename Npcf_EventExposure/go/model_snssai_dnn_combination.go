/*
 * Npcf_EventExposure
 *
 * PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_EventExposure

// SnssaiDnnCombination - Represents a combination of S-NSSAI and DNN(s).
type SnssaiDnnCombination struct {

	Snssai Snssai `json:"snssai,omitempty"`

	Dnns []string `json:"dnns,omitempty"`
}
