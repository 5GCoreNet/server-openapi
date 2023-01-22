/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsSDM

// ServiceAreaId - Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {

	PlmnId PlmnId `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac"`

	// Service Area Code.
	Sac string `json:"sac"`
}
