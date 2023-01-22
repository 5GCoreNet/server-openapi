/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

type CongestionTypeAnyOf string

// List of CongestionTypeAnyOf
const (
	USER_PLANE CongestionTypeAnyOf = "USER_PLANE"
	CONTROL_PLANE CongestionTypeAnyOf = "CONTROL_PLANE"
	USER_AND_CONTROL_PLANE CongestionTypeAnyOf = "USER_AND_CONTROL_PLANE"
)