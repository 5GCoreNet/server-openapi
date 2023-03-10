/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_Events

// TempGroupInfo - Represents the created temporary VAL group information.
type TempGroupInfo struct {

	ValGrpIds []string `json:"valGrpIds"`

	TempValGrpId string `json:"tempValGrpId"`

	ValServIds []string `json:"valServIds,omitempty"`
}
