/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_EASDiscovery

// AcProfile - ECS service provisioning response information.
type AcProfile struct {

	// Identity of the AC.
	AcId string `json:"acId"`

	// The category or type of AC.
	AcType string `json:"acType,omitempty"`

	// Indicates to the ECS which ECSPs are preferred for the AC.
	PrefEcsps []string `json:"prefEcsps,omitempty"`

	AcSchedule ScheduledCommunicationTime `json:"acSchedule,omitempty"`

	ExpAcGeoServArea LocationArea5G `json:"expAcGeoServArea,omitempty"`

	// Profiles of ACs for which the EEC provides edge enabling services.
	AcSvcContSupp []AcrScenario `json:"acSvcContSupp,omitempty"`

	// List of EAS information.
	Eass []EasDetail `json:"eass,omitempty"`
}
