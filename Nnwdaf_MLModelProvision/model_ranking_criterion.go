/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_MLModelProvision

// RankingCriterion - Indicates the usage ranking criterion between the high, medium and low usage UE.
type RankingCriterion struct {

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	HighBase int32 `json:"highBase"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	LowBase int32 `json:"lowBase"`
}
