/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// CmInfo - Represents the connection management state of a UE for an access type
type CmInfo struct {

	CmState CmState `json:"cmState"`

	AccessType AccessType `json:"accessType"`
}
