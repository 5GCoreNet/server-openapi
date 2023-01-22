/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

// DispersionRequirement - Represents the dispersion analytics requirements.
type DispersionRequirement struct {

	DisperType DispersionType `json:"disperType"`

	ClassCriters []ClassCriterion `json:"classCriters,omitempty"`

	RankCriters []RankingCriterion `json:"rankCriters,omitempty"`

	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}
