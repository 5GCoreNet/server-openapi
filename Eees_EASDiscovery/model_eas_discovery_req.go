/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EASDiscovery

// EasDiscoveryReq - ECS service provisioning request information.
type EasDiscoveryReq struct {

	RequestorId RequestorId `json:"requestorId"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId string `json:"ueId,omitempty"`

	EasDiscoveryFilter EasDiscoveryFilter `json:"easDiscoveryFilter,omitempty"`

	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EecSvcContinuity []AcrScenario `json:"eecSvcContinuity,omitempty"`

	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EesSvcContinuity []AcrScenario `json:"eesSvcContinuity,omitempty"`

	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EasSvcContinuity []AcrScenario `json:"easSvcContinuity,omitempty"`

	LocInf LocationInfo `json:"locInf,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	EasTDnai string `json:"easTDnai,omitempty"`
}
