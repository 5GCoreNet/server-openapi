/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nbsf_Management

type BsfEventAnyOf string

// List of BsfEventAnyOf
const (
	PCF_PDU_SESSION_BINDING_REGISTRATION BsfEventAnyOf = "PCF_PDU_SESSION_BINDING_REGISTRATION"
	PCF_PDU_SESSION_BINDING_DEREGISTRATION BsfEventAnyOf = "PCF_PDU_SESSION_BINDING_DEREGISTRATION"
	PCF_UE_BINDING_REGISTRATION BsfEventAnyOf = "PCF_UE_BINDING_REGISTRATION"
	PCF_UE_BINDING_DEREGISTRATION BsfEventAnyOf = "PCF_UE_BINDING_DEREGISTRATION"
	SNSSAI_DNN_BINDING_REGISTRATION BsfEventAnyOf = "SNSSAI_DNN_BINDING_REGISTRATION"
	SNSSAI_DNN_BINDING_DEREGISTRATION BsfEventAnyOf = "SNSSAI_DNN_BINDING_DEREGISTRATION"
)
