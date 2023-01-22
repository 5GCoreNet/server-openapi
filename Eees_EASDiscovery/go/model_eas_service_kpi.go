/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASDiscovery

// EasServiceKpi - Represents the EAS service KPI information.
type EasServiceKpi struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReqRate int32 `json:"maxReqRate,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxRespTime int32 `json:"maxRespTime,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Avail int32 `json:"avail,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlComp int32 `json:"avlComp,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlGraComp int32 `json:"avlGraComp,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlMem int32 `json:"avlMem,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlStrg int32 `json:"avlStrg,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConnBand string `json:"connBand,omitempty"`
}
