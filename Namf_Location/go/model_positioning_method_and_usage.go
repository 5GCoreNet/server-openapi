/*
 * Namf_Location
 *
 * AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Location

// PositioningMethodAndUsage - Indicates the usage of a positioning method.
type PositioningMethodAndUsage struct {

	Method PositioningMethod `json:"method"`

	Mode PositioningMode `json:"mode"`

	Usage Usage `json:"usage"`

	MethodCode int32 `json:"methodCode,omitempty"`
}
