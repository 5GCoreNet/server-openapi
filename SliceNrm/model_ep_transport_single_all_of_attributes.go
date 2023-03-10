/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type EpTransportSingleAllOfAttributes struct {

	IpAddress IpAddress `json:"ipAddress,omitempty"`

	LogicalInterfaceInfo LogicalInterfaceInfo `json:"logicalInterfaceInfo,omitempty"`

	NextHopInfo string `json:"nextHopInfo,omitempty"`

	QosProfile string `json:"qosProfile,omitempty"`

	EpApplicationRefs []string `json:"epApplicationRefs,omitempty"`
}
