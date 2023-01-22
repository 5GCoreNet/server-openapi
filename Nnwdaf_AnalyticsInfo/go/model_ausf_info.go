/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// AusfInfo - Information of an AUSF NF Instance
type AusfInfo struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	RoutingIndicators []string `json:"routingIndicators,omitempty"`

	SuciInfos []SuciInfo `json:"suciInfos,omitempty"`
}
