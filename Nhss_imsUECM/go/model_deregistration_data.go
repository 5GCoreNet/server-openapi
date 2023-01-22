/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsUECM

// DeregistrationData - Data related to the de-registration information of a S-CSCF in HSS
type DeregistrationData struct {

	DeregReason DeregistrationReason `json:"deregReason"`

	// IMS Private Identity of the UE
	Impi string `json:"impi"`

	AssociatedImpis []string `json:"associatedImpis,omitempty"`

	EmergencyRegisteredIdentities []EmergencyRegisteredIdentity `json:"emergencyRegisteredIdentities,omitempty"`
}
