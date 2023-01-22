/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type FrameRouteInfo struct {

	// \"String identifying a IPv4 address mask formatted in the 'dotted decimal' notation as defined in RFC 1166.\" 
	Ipv4Mask string `json:"ipv4Mask,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}
