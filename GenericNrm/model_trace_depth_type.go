/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm
// TraceDepthType : Specifies how detailed information should be recorded in the Network Element. The Trace Depth is a paremeter for Trace Session level, i.e., the Trace Depth is the same for all of the NEs to be traced in the same Trace Session. See 3GPP TS 32.422 clause 5.3 for additional details.
type TraceDepthType string

// List of TraceDepthType
const (
	MINIMUM TraceDepthType = "MINIMUM"
	MEDIUM TraceDepthType = "MEDIUM"
	MAXIMUM TraceDepthType = "MAXIMUM"
	VENDORMINIMUM TraceDepthType = "VENDORMINIMUM"
	VENDORMEDIUM TraceDepthType = "VENDORMEDIUM"
	VENDORMAXIMUM TraceDepthType = "VENDORMAXIMUM"
)
