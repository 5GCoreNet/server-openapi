/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsSDM

type RequestedNodeAnyOf string

// List of RequestedNodeAnyOf
const (
	SGSN RequestedNodeAnyOf = "SGSN"
	MME RequestedNodeAnyOf = "MME"
	AMF RequestedNodeAnyOf = "AMF"
	_3_GPP_AAA_SERVER_TWAN RequestedNodeAnyOf = "3GPP_AAA_SERVER_TWAN"
)
