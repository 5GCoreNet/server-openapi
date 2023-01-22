/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5GcNrm

type PcfFunctionSingleAllOf1 struct {

	EPN5 []EpN5Single `json:"EP_N5,omitempty"`

	EPN7 []EpN7Single `json:"EP_N7,omitempty"`

	EPN15 []EpN15Single `json:"EP_N15,omitempty"`

	EPN16 []EpN16Single `json:"EP_N16,omitempty"`

	EPRx []EpRxSingle `json:"EP_Rx,omitempty"`

	PredefinedPccRuleSet PredefinedPccRuleSetSingle `json:"PredefinedPccRuleSet,omitempty"`
}
