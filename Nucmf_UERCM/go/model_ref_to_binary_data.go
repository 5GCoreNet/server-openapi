/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nucmf_UERCM

// RefToBinaryData - This parameter provides information about the referenced binary body data.
type RefToBinaryData struct {

	// This IE shall contain the value of the Content-ID header of the referenced binary body part. 
	ContentId string `json:"contentId"`
}
