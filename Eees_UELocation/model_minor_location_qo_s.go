/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_UELocation

// MinorLocationQoS - Contain Minor Location QoS.
type MinorLocationQoS struct {

	// Indicates value of accuracy.
	HAccuracy float32 `json:"hAccuracy,omitempty"`

	// Indicates value of accuracy.
	VAccuracy float32 `json:"vAccuracy,omitempty"`
}