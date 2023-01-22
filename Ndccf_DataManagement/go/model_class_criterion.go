/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// ClassCriterion - Indicates the dispersion class criterion for fixed, camper and/or traveller UE, and/or the  top-heavy UE dispersion class criterion. 
type ClassCriterion struct {

	DisperClass DispersionClass `json:"disperClass"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	ClassThreshold int32 `json:"classThreshold"`

	ThresMatch MatchingDirection `json:"thresMatch"`
}
