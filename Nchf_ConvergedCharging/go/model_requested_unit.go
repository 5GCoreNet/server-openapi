/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

type RequestedUnit struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Time int32 `json:"time,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	TotalVolume int32 `json:"totalVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UplinkVolume int32 `json:"uplinkVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	DownlinkVolume int32 `json:"downlinkVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	ServiceSpecificUnits int32 `json:"serviceSpecificUnits,omitempty"`
}
