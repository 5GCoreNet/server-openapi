/*
 * EES Application Client Information_API
 *
 * API for EES Application Client Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_AppClientInformation

// AcFilters - Represents the filters information for AC Information Subscription.
type AcFilters struct {

	// List of AC Types or categories.
	AcTypes []string `json:"acTypes,omitempty"`

	// The list of identifiers of ECSPs associated with the EEC.
	EcspIds []string `json:"ecspIds,omitempty"`

	// List of identifiers of ACs to be matched.
	AcIds []string `json:"acIds,omitempty"`

	SvcArea ServiceArea `json:"svcArea,omitempty"`

	MaxAcKpi AcServiceKpis `json:"maxAcKpi,omitempty"`

	MinAcKpi AcServiceKpis `json:"minAcKpi,omitempty"`

	// Operation schedule of EAS to be matched with operation schedule of the AC.
	OpSchds []ScheduledCommunicationTime `json:"opSchds,omitempty"`

	// List of UE identifiers hosting the AC.
	UeIds []string `json:"ueIds,omitempty"`

	LocInfs LocationArea5G `json:"locInfs,omitempty"`
}
