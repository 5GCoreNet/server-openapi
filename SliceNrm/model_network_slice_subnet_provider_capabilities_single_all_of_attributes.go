/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes struct {

	DLlatency int32 `json:"dLlatency,omitempty"`

	ULlatency int32 `json:"uLlatency,omitempty"`

	DLThptPerSliceSubnet XlThpt `json:"dLThptPerSliceSubnet,omitempty"`

	ULThptPerSliceSubnet XlThpt `json:"uLThptPerSliceSubnet,omitempty"`

	CoverageAreaTAList []int32 `json:"coverageAreaTAList,omitempty"`
}
