/*
 * M1_EdgeResourcesProvisioning
 *
 * 5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package M1_EdgeResourcesProvisioning

type PointAltitudeAllOf struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}
