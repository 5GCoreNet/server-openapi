/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

type MediaStreamingAccessRecordAllOfResponseMessage struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ResponseCode int32 `json:"responseCode"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Size int32 `json:"size"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BodySize int32 `json:"bodySize"`

	ContentType string `json:"contentType,omitempty"`
}
