/*
 * 3gpp-eas-deployment
 *
 * API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EASDeployment

// EdiDeleteCriteria - Contains criteria to be used for deleting EAS Deployment Information entries that match them.
type EdiDeleteCriteria struct {

	// Represents an AF identifier.
	AfId string `json:"afId,omitempty"`

	DnnSnssai DnnSnssaiInformation `json:"dnnSnssai,omitempty"`
}
