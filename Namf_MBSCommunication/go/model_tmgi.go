/*
 * Namf_MBSCommunication
 *
 * AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_MBSCommunication

// Tmgi - Temporary Mobile Group Identity
type Tmgi struct {

	// MBS Service ID
	MbsServiceId string `json:"mbsServiceId"`

	PlmnId PlmnId `json:"plmnId"`
}
