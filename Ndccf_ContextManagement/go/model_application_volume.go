/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// ApplicationVolume - Application data volume per Application Id.
type ApplicationVolume struct {

	// String providing an application identifier.
	AppId string `json:"appId"`

	// Unsigned integer identifying a volume in units of bytes.
	AppVolume int64 `json:"appVolume"`
}
