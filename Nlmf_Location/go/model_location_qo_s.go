/*
 * LMF Location
 *
 * LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nlmf_Location

// LocationQoS - QoS of Location request.
type LocationQoS struct {

	// Indicates value of accuracy.
	HAccuracy float32 `json:"hAccuracy,omitempty"`

	// Indicates value of accuracy.
	VAccuracy float32 `json:"vAccuracy,omitempty"`

	VerticalRequested bool `json:"verticalRequested,omitempty"`

	ResponseTime ResponseTime `json:"responseTime,omitempty"`

	MinorLocQoses []MinorLocationQoS `json:"minorLocQoses,omitempty"`

	LcsQosClass LcsQosClass `json:"lcsQosClass,omitempty"`
}
