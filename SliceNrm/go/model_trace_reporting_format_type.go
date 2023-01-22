/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm
// TraceReportingFormatType : Specifies whether file-based or streaming reporting shall be used for this Trace Session. See 3GPP TS 32.422 clause 5.11 for additional details.
type TraceReportingFormatType string

// List of TraceReportingFormatType
const (
	FILE_BASED TraceReportingFormatType = "FILE-BASED"
	STREAMING TraceReportingFormatType = "STREAMING"
)
