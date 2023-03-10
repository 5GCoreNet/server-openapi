/*
 * UAE Server C2 Operation Mode Management Service
 *
 * UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_C2OperationModeManagement

// C2OpModeMngtCompStatus - Represents the C2 Operation Mode Management Completion status for a UAV (e.g. UAV, UAV-C). 
type C2OpModeMngtCompStatus struct {

	UasId UasId `json:"uasId"`

	Status C2OpModeStatus `json:"status"`
}
