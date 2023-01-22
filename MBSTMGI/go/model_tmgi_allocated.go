/*
 * 3gpp-mbs-tmgi
 *
 * API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSTMGI

import (
	"time"
)

// TmgiAllocated - Data within TMGI Allocate Response
type TmgiAllocated struct {

	// One or more TMGI values
	TmgiList []Tmgi `json:"tmgiList"`

	// string with format 'date-time' as defined in OpenAPI.
	ExpirationTime time.Time `json:"expirationTime"`
}
