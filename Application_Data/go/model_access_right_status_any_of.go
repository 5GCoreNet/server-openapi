/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Application_Data

type AccessRightStatusAnyOf string

// List of AccessRightStatusAnyOf
const (
	FULLY_ALLOWED AccessRightStatusAnyOf = "FULLY_ALLOWED"
	PREVIEW_ALLOWED AccessRightStatusAnyOf = "PREVIEW_ALLOWED"
	NO_ALLOWED AccessRightStatusAnyOf = "NO_ALLOWED"
)
