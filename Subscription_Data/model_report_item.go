/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// ReportItem - indicates performed modivications.
type ReportItem struct {

	// Contains a JSON pointer value (as defined in IETF RFC 6901) that references a  location of a resource to which the modification is subject. 
	Path string `json:"path"`

	// A human-readable reason providing details on the reported modification failure.  The reason string should identify the operation that failed using the operation's  array index to assist in correlation of the invalid parameter with the failed  operation, e.g. \"Replacement value invalid for attribute (failed operation index= 4)\". 
	Reason string `json:"reason,omitempty"`
}
