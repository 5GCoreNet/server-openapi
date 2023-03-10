/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

type ServAuthInfoAnyOf string

// List of ServAuthInfoAnyOf
const (
	TP_NOT_KNOWN ServAuthInfoAnyOf = "TP_NOT_KNOWN"
	TP_EXPIRED ServAuthInfoAnyOf = "TP_EXPIRED"
	TP_NOT_YET_OCURRED ServAuthInfoAnyOf = "TP_NOT_YET_OCURRED"
	ROUT_REQ_NOT_AUTHORIZED ServAuthInfoAnyOf = "ROUT_REQ_NOT_AUTHORIZED"
)
