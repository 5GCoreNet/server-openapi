/*
 * Neasdf_DNSContext
 *
 * EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Neasdf_DNSContext

// StringMatchingCondition - A String with Matching Operator
type StringMatchingCondition struct {

	MatchingString string `json:"matchingString,omitempty"`

	MatchingOperator MatchingOperator `json:"matchingOperator"`
}
