/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_PP

// Tmgi - Temporary Mobile Group Identity
type Tmgi struct {

	// MBS Service ID
	MbsServiceId string `json:"mbsServiceId"`

	PlmnId PlmnId `json:"plmnId"`
}
