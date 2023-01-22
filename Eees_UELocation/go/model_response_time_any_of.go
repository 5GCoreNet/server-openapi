/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_UELocation

type ResponseTimeAnyOf string

// List of ResponseTimeAnyOf
const (
	LOW_DELAY ResponseTimeAnyOf = "LOW_DELAY"
	DELAY_TOLERANT ResponseTimeAnyOf = "DELAY_TOLERANT"
	NO_DELAY ResponseTimeAnyOf = "NO_DELAY"
)
