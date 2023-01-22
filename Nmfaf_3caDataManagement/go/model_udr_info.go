/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// UdrInfo - Information of an UDR NF Instance
type UdrInfo struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	SupportedDataSets []DataSetId `json:"supportedDataSets,omitempty"`

	SharedDataIdRanges []SharedDataIdRange `json:"sharedDataIdRanges,omitempty"`
}
