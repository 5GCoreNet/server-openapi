/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// CommunicationFailure - Describes a communication failure detected by AMF
type CommunicationFailure struct {

	NasReleaseCode string `json:"nasReleaseCode,omitempty"`

	RanReleaseCode NgApCause `json:"ranReleaseCode,omitempty"`
}
