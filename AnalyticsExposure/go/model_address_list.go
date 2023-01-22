/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

// AddressList - Represents a list of IPv4 and/or IPv6 addresses.
type AddressList struct {

	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`
}
