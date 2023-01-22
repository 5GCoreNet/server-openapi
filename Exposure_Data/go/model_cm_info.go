/*
 * Unified Data Repository Service API file for structured data for exposure
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Exposure_Data

// CmInfo - Represents the connection management state of a UE for an access type
type CmInfo struct {

	CmState CmState `json:"cmState"`

	AccessType AccessType `json:"accessType"`
}
