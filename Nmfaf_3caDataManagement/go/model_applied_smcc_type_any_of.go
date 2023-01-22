/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement
// AppliedSmccTypeAnyOf : This string indicates the type of applied SM congestion control. 
type AppliedSmccTypeAnyOf string

// List of AppliedSmccTypeAnyOf
const (
	DNN_CC AppliedSmccTypeAnyOf = "DNN_CC"
	SNSSAI_CC AppliedSmccTypeAnyOf = "SNSSAI_CC"
)