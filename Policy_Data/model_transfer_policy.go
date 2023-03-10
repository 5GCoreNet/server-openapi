/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

// TransferPolicy - Describes a transfer policy.
type TransferPolicy struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBitRateDl string `json:"maxBitRateDl,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBitRateUl string `json:"maxBitRateUl,omitempty"`

	// Indicates a rating group for the recommended time window.
	RatingGroup int32 `json:"ratingGroup"`

	RecTimeInt TimeWindow `json:"recTimeInt"`

	// Contains an identity of a transfer policy.
	TransPolicyId int32 `json:"transPolicyId"`
}
