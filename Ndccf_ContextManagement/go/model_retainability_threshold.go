/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// RetainabilityThreshold - Represents a QoS flow retainability threshold.
type RetainabilityThreshold struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelFlowNum int32 `json:"relFlowNum,omitempty"`

	RelTimeUnit TimeUnit `json:"relTimeUnit,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	RelFlowRatio int32 `json:"relFlowRatio,omitempty"`
}
