/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

// AfRoutingRequirementRm - This data type is defined in the same way as the AfRoutingRequirement data type, but with the OpenAPI nullable property set to true and the spVal and tempVals attributes defined as removable. 
type AfRoutingRequirementRm struct {

	AppReloc bool `json:"appReloc,omitempty"`

	RouteToLocs *[]RouteToLocation `json:"routeToLocs,omitempty"`

	SpVal *SpatialValidityRm `json:"spVal,omitempty"`

	TempVals *[]TemporalValidity `json:"tempVals,omitempty"`

	UpPathChgSub *UpPathChgEvent `json:"upPathChgSub,omitempty"`

	AddrPreserInd *bool `json:"addrPreserInd,omitempty"`

	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd *bool `json:"simConnInd,omitempty"`

	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	SimConnTerm *int32 `json:"simConnTerm,omitempty"`

	// Contains EAS IP replacement information.
	EasIpReplaceInfos *[]EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`

	// Indicates the EAS rediscovery is required.
	EasRedisInd bool `json:"easRedisInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	MaxAllowedUpLat *int32 `json:"maxAllowedUpLat,omitempty"`
}
