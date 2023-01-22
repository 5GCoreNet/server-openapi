/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

// SessionRule - Contains session level policy information.
type SessionRule struct {

	AuthSessAmbr Ambr `json:"authSessAmbr,omitempty"`

	AuthDefQos AuthorizedDefaultQos `json:"authDefQos,omitempty"`

	// Univocally identifies the session rule within a PDU session.
	SessRuleId string `json:"sessRuleId"`

	// A reference to UsageMonitoringData policy decision type. It is the umId described in clause 5.6.2.12.
	RefUmData *string `json:"refUmData,omitempty"`

	// A reference to UsageMonitoringData policy decision type to apply for Non-3GPP access. It is the umId described in clause 5.6.2.12.
	RefUmN3gData *string `json:"refUmN3gData,omitempty"`

	// A reference to the condition data. It is the condId described in clause 5.6.2.9.
	RefCondData *string `json:"refCondData,omitempty"`
}
