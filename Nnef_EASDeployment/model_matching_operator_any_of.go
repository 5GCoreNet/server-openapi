/*
 * Nnef_EASDeployment
 *
 * NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnef_EASDeployment

type MatchingOperatorAnyOf string

// List of MatchingOperatorAnyOf
const (
	FULL_MATCH MatchingOperatorAnyOf = "FULL_MATCH"
	MATCH_ALL MatchingOperatorAnyOf = "MATCH_ALL"
	STARTS_WITH MatchingOperatorAnyOf = "STARTS_WITH"
	NOT_START_WITH MatchingOperatorAnyOf = "NOT_START_WITH"
	ENDS_WITH MatchingOperatorAnyOf = "ENDS_WITH"
	NOT_END_WITH MatchingOperatorAnyOf = "NOT_END_WITH"
	CONTAINS MatchingOperatorAnyOf = "CONTAINS"
	NOT_CONTAIN MatchingOperatorAnyOf = "NOT_CONTAIN"
)
