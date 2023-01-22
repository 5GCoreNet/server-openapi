/*
 * Eecs_ServiceProvisioning
 *
 * API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eecs_ServiceProvisioning

// EcsServProvReq - ECS service provisioning request information.
type EcsServProvReq struct {

	// Represents a unique identifier of the EEC.
	EecId string `json:"eecId"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	UeId string `json:"ueId,omitempty"`

	// Information about services the EEC wants to connect to.
	AcProfs []AcProfile `json:"acProfs,omitempty"`

	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. 
	EecSvcContSupp []AcrScenario `json:"eecSvcContSupp,omitempty"`

	// List of connectivity information for the UE.
	ConnInfo []ConnectivityInfo `json:"connInfo,omitempty"`

	LocInf LocationInfo `json:"locInf,omitempty"`
}
