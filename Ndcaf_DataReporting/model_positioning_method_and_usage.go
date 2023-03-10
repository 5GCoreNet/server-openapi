/*
 * Ndcaf_DataReporting
 *
 * Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndcaf_DataReporting

// PositioningMethodAndUsage - Indicates the usage of a positioning method.
type PositioningMethodAndUsage struct {

	Method PositioningMethod `json:"method"`

	Mode PositioningMode `json:"mode"`

	Usage Usage `json:"usage"`

	MethodCode int32 `json:"methodCode,omitempty"`
}
