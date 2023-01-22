/*
 * 5GMS Common Data Types
 *
 * 5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

type ChargingSpecification struct {

	SponId string `json:"sponId,omitempty"`

	SponStatus SponsoringStatus `json:"sponStatus,omitempty"`

	Gpsi []string `json:"gpsi,omitempty"`
}
