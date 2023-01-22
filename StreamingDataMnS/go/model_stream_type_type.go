/*
 * TS 28.532 Streaming data reporting service
 *
 * OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package StreamingDataMnS

type StreamTypeType string

// List of StreamTypeType
const (
	TRACE StreamTypeType = "TRACE"
	PERFORMANCE StreamTypeType = "PERFORMANCE"
	ANALYTICS StreamTypeType = "ANALYTICS"
	PROPRIETARY StreamTypeType = "PROPRIETARY"
)
