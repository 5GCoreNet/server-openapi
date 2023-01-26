/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSSession

type MbsSessionEventTypeAnyOf string

// List of MbsSessionEventTypeAnyOf
const (
	MBS_REL_TMGI_EXPIRY MbsSessionEventTypeAnyOf = "MBS_REL_TMGI_EXPIRY"
	BROADCAST_DELIVERY_STATUS MbsSessionEventTypeAnyOf = "BROADCAST_DELIVERY_STATUS"
	INGRESS_TUNNEL_ADD_CHANGE MbsSessionEventTypeAnyOf = "INGRESS_TUNNEL_ADD_CHANGE"
)