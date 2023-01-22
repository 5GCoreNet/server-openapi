/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_AMPolicyControl

// TerminationNotification - Represents a request to terminate a policy Association that the PCF provides in a notification. 
type TerminationNotification struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`

	Cause PolicyAssociationReleaseCause `json:"cause"`
}
