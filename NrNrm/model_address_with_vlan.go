/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type AddressWithVlan struct {

	Ipv4Address string `json:"ipv4Address,omitempty"`

	Ipv6Address Ipv6Addr `json:"ipv6Address,omitempty"`

	VlanId int32 `json:"vlanId,omitempty"`
}
