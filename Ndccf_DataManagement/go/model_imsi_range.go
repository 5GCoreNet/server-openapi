/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// ImsiRange - A range of IMSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type ImsiRange struct {

	Start string `json:"start,omitempty"`

	End string `json:"end,omitempty"`

	Pattern string `json:"pattern,omitempty"`
}
