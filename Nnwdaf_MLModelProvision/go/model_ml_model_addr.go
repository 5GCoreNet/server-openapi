/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_MLModelProvision

// MlModelAddr - Addresses of ML model files.
type MlModelAddr struct {

	// String providing an URI formatted according to RFC 3986.
	MLModelUrl string `json:"mLModelUrl,omitempty"`

	// The FQDN of the ML Model file.
	MlFileFqdn string `json:"mlFileFqdn,omitempty"`
}
