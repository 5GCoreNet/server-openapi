/*
 * GBA BSF Nbsp_GBA Service
 *
 * GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nbsp_GBA

import (
	"time"
)

// BootstrappingInfoResponse - Response body of the HTTP POST operation for resource /bootstrapping-info-request
type BootstrappingInfoResponse struct {

	// ME Key Material (hex-encoded string)
	MeKeyMaterial string `json:"meKeyMaterial"`

	// UICC key material (hex-encoded string)
	UiccKeyMaterial string `json:"uiccKeyMaterial,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	KeyExpiryTime time.Time `json:"keyExpiryTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	BootstrappingInfoCreationTime time.Time `json:"bootstrappingInfoCreationTime,omitempty"`

	UssList []UssListItem `json:"ussList,omitempty"`

	GbaType GbaType `json:"gbaType,omitempty"`

	// IMS Private Identity of the UE
	Impi string `json:"impi,omitempty"`
}
