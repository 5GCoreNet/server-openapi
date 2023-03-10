/*
 * Nnsacf_NSAC
 *
 * Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnsacf_NSAC

type PduAcRequestData struct {

	PduACRequestInfo []PduAcRequestInfo `json:"pduACRequestInfo"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfId string `json:"nfId,omitempty"`

	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn,omitempty"`
}
