/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// QosFlowProfile - MBS QoS flow profile
type QosFlowProfile struct {

	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi"`

	NonDynamic5Qi NonDynamic5Qi `json:"nonDynamic5Qi,omitempty"`

	Dynamic5Qi Dynamic5Qi `json:"dynamic5Qi,omitempty"`

	Arp Arp `json:"arp,omitempty"`

	GbrQosFlowInfo GbrQosFlowInformation `json:"gbrQosFlowInfo,omitempty"`
}