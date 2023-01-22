/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// AfEventExposureData - AF Event Exposure data managed by a given NEF Instance
type AfEventExposureData struct {

	AfEvents []AfEvent `json:"afEvents"`

	AfIds []string `json:"afIds,omitempty"`

	AppIds []string `json:"appIds,omitempty"`
}
