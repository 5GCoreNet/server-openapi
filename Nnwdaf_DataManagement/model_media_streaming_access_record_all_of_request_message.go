/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

type MediaStreamingAccessRecordAllOfRequestMessage struct {

	Method string `json:"method"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	Url string `json:"url"`

	ProtocolVersion string `json:"protocolVersion"`

	Range string `json:"range,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Size int32 `json:"size"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BodySize int32 `json:"bodySize"`

	ContentType string `json:"contentType,omitempty"`

	UserAgent string `json:"userAgent,omitempty"`

	UserIdentity string `json:"userIdentity,omitempty"`

	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	Referer string `json:"referer,omitempty"`
}
