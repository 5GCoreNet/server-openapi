/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

// EasServerAddress - Represents the IP address and port of an EAS server.
type EasServerAddress struct {

	Ip IpAddr `json:"ip"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Port int32 `json:"port"`
}
