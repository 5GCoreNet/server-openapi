/*
 * 3gpp-mbs-tmgi
 *
 * API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSTMGI

// TmgiDeallocRequest - Represents information to request the deallocation of MBS TMGI(s).
type TmgiDeallocRequest struct {

	AfId string `json:"afId"`

	Tmgis []Tmgi `json:"tmgis"`
}
