/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type MaxNumberofUes struct {

	ServAttrCom ServAttrCom `json:"servAttrCom,omitempty"`

	Var3GPPNoOfUEs int32 `json:"3GPPNoOfUEs,omitempty"`

	Non3GPPNoOfUEs int32 `json:"non3GPPNoOfUEs,omitempty"`
}