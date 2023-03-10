/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

type NrfInfoServedHssInfoListValueValue struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	ImsiRanges []ImsiRange `json:"imsiRanges,omitempty"`

	ImsPrivateIdentityRanges []IdentityRange `json:"imsPrivateIdentityRanges,omitempty"`

	ImsPublicIdentityRanges []IdentityRange `json:"imsPublicIdentityRanges,omitempty"`

	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	HssDiameterAddress NetworkNodeDiameterAddress `json:"hssDiameterAddress,omitempty"`

	AdditionalDiamAddresses []NetworkNodeDiameterAddress `json:"additionalDiamAddresses,omitempty"`
}
