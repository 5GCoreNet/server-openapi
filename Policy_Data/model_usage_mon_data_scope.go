/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

// UsageMonDataScope - Contains a SNSSAI and DNN combinations to which the UsageMonData instance belongs to. 
type UsageMonDataScope struct {

	Snssai Snssai `json:"snssai"`

	Dnn []string `json:"dnn,omitempty"`
}
