/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AiMlNrm

// ReportIntervalType - See details in 3GPP TS 32.422 clause 5.10.5.
type ReportIntervalType struct {

	UMTS []string `json:"UMTS,omitempty"`

	LTE []string `json:"LTE,omitempty"`

	NR []string `json:"NR,omitempty"`
}
