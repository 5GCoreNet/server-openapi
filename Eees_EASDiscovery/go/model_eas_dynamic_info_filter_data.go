/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASDiscovery

// EasDynamicInfoFilterData - Represents an EAS dynamic information.
type EasDynamicInfoFilterData struct {

	// Represents a unique identifier of the EEC.
	EecId string `json:"eecId"`

	// Notify if EAS status changed.
	EasStatus bool `json:"easStatus,omitempty"`

	// Notify if list of AC identifiers changed.
	EasAcIds bool `json:"easAcIds,omitempty"`

	// Notify if EAS description changed.
	EasDesc bool `json:"easDesc,omitempty"`

	// Notify if EAS endpoint changed.
	EasPt bool `json:"easPt,omitempty"`

	// NotiNotify if EAS feature changed.
	EasFeature bool `json:"easFeature,omitempty"`

	// Notify if EAS schedule changed.
	EasSchedule bool `json:"easSchedule,omitempty"`

	// Notify if EAS service area changed.
	SvcArea bool `json:"svcArea,omitempty"`

	// Notify if EAS KPIs changed.
	SvcKpi bool `json:"svcKpi,omitempty"`

	// Notify if EAS supported ACR changed.
	SvcCont bool `json:"svcCont,omitempty"`
}
