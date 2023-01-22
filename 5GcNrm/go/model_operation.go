/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5GcNrm

type Operation struct {

	Name string `json:"name,omitempty"`

	AllowedNFTypes NfType `json:"allowedNFTypes,omitempty"`

	OperationSemantics OperationSemantics `json:"operationSemantics,omitempty"`
}
