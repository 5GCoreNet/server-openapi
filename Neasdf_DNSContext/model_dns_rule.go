/*
 * Neasdf_DNSContext
 *
 * EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Neasdf_DNSContext

// DnsRule - DNS message handling rule
type DnsRule struct {

	DnsRuleId string `json:"dnsRuleId,omitempty"`

	Label string `json:"label,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Precedence int32 `json:"precedence,omitempty"`

	// map of DNS query message detection templates where a valid JSON string serves as key
	DnsQueryMdtList map[string]DnsQueryMdt `json:"dnsQueryMdtList,omitempty"`

	BaseDnsQueryMdtList []BaselineDnsQueryMdtInfo `json:"baseDnsQueryMdtList,omitempty"`

	// map of DNS response message detection templates where a valid JSON string serves as key
	DnsRspMdtList map[string]DnsRspMdt `json:"dnsRspMdtList,omitempty"`

	BaseDnsRspMdtList []BaselineDnsRspMdtInfo `json:"baseDnsRspMdtList,omitempty"`

	DnsMsgId string `json:"dnsMsgId,omitempty"`

	// map of actions where a valid JSON string serves as key
	ActionList map[string]Action `json:"actionList"`
}
