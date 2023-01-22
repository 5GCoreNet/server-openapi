/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CoslaNrm
// JobTypeType : Specifies whether the TraceJob represents only MDT, Logged MBSFN MDT, Trace or a combined Trace and MDT job. Applicable for Trace, MDT, RCEF and RLF reporting. See 3GPP TS 32.422 clause 5.9a for additional details.
type JobTypeType string

// List of JobTypeType
const (
	IMMEDIATE_MDT_ONLY JobTypeType = "IMMEDIATE_MDT_ONLY"
	LOGGED_MDT_ONLY JobTypeType = "LOGGED_MDT_ONLY"
	TRACE_ONLY JobTypeType = "TRACE_ONLY"
	IMMEDIATE_MDT_AND_TRACE JobTypeType = "IMMEDIATE_MDT AND TRACE"
	RLF_REPORT_ONLY JobTypeType = "RLF_REPORT_ONLY"
	RCEF_REPORT_ONLY JobTypeType = "RCEF_REPORT_ONLY"
	LOGGED_MBSFN_MDT JobTypeType = "LOGGED_MBSFN_MDT"
)
