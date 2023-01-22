/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// FqdnPatternMatchingRule - a matching rule for a FQDN pattern
type FqdnPatternMatchingRule struct {

	Regex string `json:"regex,omitempty"`

	StringMatchingRule StringMatchingRule `json:"stringMatchingRule,omitempty"`
}
