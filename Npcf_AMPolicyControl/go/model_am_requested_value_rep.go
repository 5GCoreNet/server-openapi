/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_AMPolicyControl

// AmRequestedValueRep - Represents the current applicable values corresponding to the policy control request triggers. 
type AmRequestedValueRep struct {

	UserLoc UserLocation `json:"userLoc,omitempty"`

	// Contains the UE presence statuses for tracking areas. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraStatuses map[string]PresenceInfo `json:"praStatuses,omitempty"`

	AccessTypes []AccessType `json:"accessTypes,omitempty"`

	RatTypes []RatType `json:"ratTypes,omitempty"`

	// array of allowed S-NSSAIs for the 3GPP access.
	AllowedSnssais []Snssai `json:"allowedSnssais,omitempty"`

	// array of allowed S-NSSAIs for the Non-3GPP access.
	N3gAllowedSnssais []Snssai `json:"n3gAllowedSnssais,omitempty"`
}
