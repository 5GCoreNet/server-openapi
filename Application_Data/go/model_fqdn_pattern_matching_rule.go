/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Application_Data

// FqdnPatternMatchingRule - a matching rule for a FQDN pattern
type FqdnPatternMatchingRule struct {

	Regex string `json:"regex,omitempty"`

	StringMatchingRule StringMatchingRule `json:"stringMatchingRule,omitempty"`
}
