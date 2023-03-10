/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

// V2xContext - Represents the V2X services related parameters
type V2xContext struct {

	NrV2xServicesAuth NrV2xAuth `json:"nrV2xServicesAuth,omitempty"`

	LteV2xServicesAuth LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	NrUeSidelinkAmbr string `json:"nrUeSidelinkAmbr,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	LteUeSidelinkAmbr string `json:"lteUeSidelinkAmbr,omitempty"`

	Pc5QoSPara Pc5QoSPara `json:"pc5QoSPara,omitempty"`
}
