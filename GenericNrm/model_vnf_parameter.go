/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

type VnfParameter struct {

	VnfInstanceId string `json:"vnfInstanceId,omitempty"`

	VnfdId string `json:"vnfdId,omitempty"`

	FlavourId string `json:"flavourId,omitempty"`

	AutoScalable bool `json:"autoScalable,omitempty"`
}
