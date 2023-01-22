/*
 * JOSE Protected Message Forwarding API
 *
 * N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package JOSEProtectedMessageForwarding

// IndexToEncryptedValue - Index to the encrypted value
type IndexToEncryptedValue struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	EncBlockIndex int32 `json:"encBlockIndex"`
}
