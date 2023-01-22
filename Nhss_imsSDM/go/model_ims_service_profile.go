/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsSDM

// ImsServiceProfile - IMS Service Profile of the UE, containing the list of Public Identifiers and optionally a list of IFCs 
type ImsServiceProfile struct {

	PublicIdentifierList []PublicIdentifier `json:"publicIdentifierList"`

	Ifcs Ifcs `json:"ifcs,omitempty"`

	CnServiceAuthorization CoreNetworkServiceAuthorization `json:"cnServiceAuthorization,omitempty"`
}
