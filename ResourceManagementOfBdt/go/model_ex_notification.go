/*
 * 3gpp-bdt
 *
 * API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ResourceManagementOfBdt

// ExNotification - Represents a Background Data Transfer notification.
type ExNotification struct {

	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`

	LocationArea5G LocationArea5G `json:"locationArea5G,omitempty"`

	TimeWindow TimeWindow `json:"timeWindow,omitempty"`

	// This IE indicates a list of the candidate transfer policies from which the AF may select a new transfer policy due to network performance degradation.
	CandPolicies []TransferPolicy `json:"candPolicies,omitempty"`
}
