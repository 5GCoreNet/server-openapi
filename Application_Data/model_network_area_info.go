/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Application_Data

// NetworkAreaInfo - Describes a network area information in which the NF service consumer requests the number of UEs. 
type NetworkAreaInfo struct {

	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`

	// Contains a list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`

	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`

	// Contains a list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`
}
