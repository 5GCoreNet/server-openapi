/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

// NetworkAreaInfo1 - Describes a network area information in which the NF service consumer requests the number of UEs.
type NetworkAreaInfo1 struct {

	// Contains a list of E-UTRA cell identities.
	Ecgis []Ecgi `json:"ecgis,omitempty"`

	// Contains a list of NR cell identities.
	Ncgis []Ncgi `json:"ncgis,omitempty"`

	// Contains a list of NG RAN nodes.
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`

	// Contains a list of tracking area identities.
	Tais []Tai `json:"tais,omitempty"`
}
