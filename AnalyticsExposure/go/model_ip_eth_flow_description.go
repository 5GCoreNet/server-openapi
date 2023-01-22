/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

// IpEthFlowDescription - Contains the description of an Uplink and/or Downlink Ethernet flow.
type IpEthFlowDescription struct {

	// Defines a packet filter of an IP flow.
	IpTrafficFilter string `json:"ipTrafficFilter,omitempty"`

	EthTrafficFilter EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}
