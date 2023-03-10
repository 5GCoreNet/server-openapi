/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_N5g-ddnmf_Discovery

// MonitorAuthReqData - Represents Data used to request the authorization to monitor for a UE
type MonitorAuthReqData struct {

	DiscType DiscoveryType `json:"discType"`

	OpenDiscData MonitorDiscDataForOpen `json:"openDiscData,omitempty"`

	RestrictedDiscData MonitorDiscDataForRestricted `json:"restrictedDiscData,omitempty"`
}
