/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nucmf_UERCM

import (
	"os"
)

type CreateDictionaryEntryRequest struct {

	JsonData DicEntryCreateData `json:"jsonData,omitempty"`

	BinaryDataUeRadioCapability5GS *os.File `json:"binaryDataUeRadioCapability5GS,omitempty"`

	BinaryDataUeRadioCapabilityEPS *os.File `json:"binaryDataUeRadioCapabilityEPS,omitempty"`

	BinaryDataUeRadioCap5GSForPaging *os.File `json:"binaryDataUeRadioCap5GSForPaging,omitempty"`

	BinaryDataUeRadioCapEPSForPaging *os.File `json:"binaryDataUeRadioCapEPSForPaging,omitempty"`
}
