/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

type PeeParameter struct {

	SiteIdentification string `json:"siteIdentification,omitempty"`

	SiteDescription string `json:"siteDescription,omitempty"`

	SiteLatitude float32 `json:"siteLatitude,omitempty"`

	SiteLongitude float32 `json:"siteLongitude,omitempty"`

	SiteAltitude float32 `json:"siteAltitude,omitempty"`

	EquipmentType string `json:"equipmentType,omitempty"`

	EnvironmentType string `json:"environmentType,omitempty"`

	PowerInterface string `json:"powerInterface,omitempty"`
}
