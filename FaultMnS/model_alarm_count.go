/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_FaultMnS

type AlarmCount struct {

	CriticalCount int32 `json:"criticalCount"`

	MajorCount int32 `json:"majorCount"`

	MinorCount int32 `json:"minorCount"`

	WarningCount int32 `json:"warningCount"`

	IndeterminateCount int32 `json:"indeterminateCount"`

	ClearedCount int32 `json:"clearedCount"`
}
