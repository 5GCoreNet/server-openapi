/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

type RemoteAddress struct {

	Ipv4Address string `json:"ipv4Address,omitempty"`

	Ipv6Address Ipv6Addr `json:"ipv6Address,omitempty"`
}