/*
 * Nmfaf_3daDataManagement
 *
 * MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3daDataManagement

// MfafNotiInfo - The MFAF notification information. It shall be provided in a response message if it had not been provided in the respective request message. 
type MfafNotiInfo struct {

	// String providing an URI formatted according to RFC 3986.
	MfafNotifUri string `json:"mfafNotifUri"`

	MfafCorreId string `json:"mfafCorreId"`
}
