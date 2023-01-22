/*
 * Nnsacf_SliceEventExposure
 *
 * Nnsacf_SliceEventExposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnsacf_SliceEventExposure

// PatchItem - it contains information on data to be changed.
type PatchItem struct {

	Op PatchOperation `json:"op"`

	// contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of a resource on which the patch operation shall be performed. 
	Path string `json:"path"`

	// indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the \"path\" attribute. 
	From string `json:"from,omitempty"`

	Value *interface{} `json:"value,omitempty"`
}
