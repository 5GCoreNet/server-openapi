/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

type SharedDataTreatmentInstructionAnyOf string

// List of SharedDataTreatmentInstructionAnyOf
const (
	USE_IF_NO_CLASH SharedDataTreatmentInstructionAnyOf = "USE_IF_NO_CLASH"
	OVERWRITE SharedDataTreatmentInstructionAnyOf = "OVERWRITE"
	MAX SharedDataTreatmentInstructionAnyOf = "MAX"
	MIN SharedDataTreatmentInstructionAnyOf = "MIN"
)
