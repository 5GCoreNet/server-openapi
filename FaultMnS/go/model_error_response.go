/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package FaultMnS

// ErrorResponse - Default schema for the response message body in case the request is not successful.
type ErrorResponse struct {

	Error ErrorResponseError `json:"error,omitempty"`
}
