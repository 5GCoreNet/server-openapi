/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5GcNrm

type IpInterface struct {

	Ipv4EndpointAddresses string `json:"ipv4EndpointAddresses,omitempty"`

	Ipv6EndpointAddresses Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`

	Fqdn string `json:"fqdn,omitempty"`
}
