/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_PDUSession

type N4MessageTypeAnyOf string

// List of N4MessageTypeAnyOf
const (
	EST_REQ N4MessageTypeAnyOf = "PFCP_SES_EST_REQ"
	EST_RSP N4MessageTypeAnyOf = "PFCP_SES_EST_RSP"
	MOD_REQ N4MessageTypeAnyOf = "PFCP_SES_MOD_REQ"
	MOD_RSP N4MessageTypeAnyOf = "PFCP_SES_MOD_RSP"
	DEL_REQ N4MessageTypeAnyOf = "PFCP_SES_DEL_REQ"
	DEL_RSP N4MessageTypeAnyOf = "PFCP_SES_DEL_RSP"
	REP_REQ N4MessageTypeAnyOf = "PFCP_SES_REP_REQ"
	REP_RSP N4MessageTypeAnyOf = "PFCP_SES_REP_RSP"
)
