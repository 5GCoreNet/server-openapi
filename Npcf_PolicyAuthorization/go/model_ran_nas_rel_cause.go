/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

// RanNasRelCause - Contains the RAN/NAS release cause.
type RanNasRelCause struct {

	NgApCause NgApCause `json:"ngApCause,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gMmCause int32 `json:"5gMmCause,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Var5gSmCause int32 `json:"5gSmCause,omitempty"`

	// Defines the EPS RAN/NAS release cause.
	EpsCause string `json:"epsCause,omitempty"`
}
