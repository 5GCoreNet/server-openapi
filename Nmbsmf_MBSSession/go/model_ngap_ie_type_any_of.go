/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsmf_MBSSession

type NgapIeTypeAnyOf string

// List of NgapIeTypeAnyOf
const (
	SETUP_REQ NgapIeTypeAnyOf = "MBS_DIS_SETUP_REQ"
	SETUP_RSP NgapIeTypeAnyOf = "MBS_DIS_SETUP_RSP"
	SETUP_FAIL NgapIeTypeAnyOf = "MBS_DIS_SETUP_FAIL"
	REL_REQ NgapIeTypeAnyOf = "MBS_DIS_REL_REQ"
)
