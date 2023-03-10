/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// EasServerAddress - Represents the IP address and port of an EAS server.
type EasServerAddress struct {

	Ip IpAddr `json:"ip"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Port int32 `json:"port"`
}
