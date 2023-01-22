/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

// EasIpReplacementInfo - Contains EAS IP replacement information for a Source and a Target EAS.
type EasIpReplacementInfo struct {

	Source EasServerAddress `json:"source"`

	Target EasServerAddress `json:"target"`
}
