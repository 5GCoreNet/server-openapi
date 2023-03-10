/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// FailureEventInfo - Contains information on the event for which the subscription is not successful.
type FailureEventInfo struct {

	Event NwdafEvent `json:"event"`

	FailureCode NwdafFailureCode `json:"failureCode"`
}
