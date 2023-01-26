/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

// TraceTargetType - Trace target conveying both the type and value of the target ID. For additional details see 3GPP TS 32.422
type TraceTargetType struct {

	TargetIdType string `json:"TargetIdType"`

	TargetIdValue string `json:"TargetIdValue"`
}