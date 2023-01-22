/*
 * Nsoraf_SOR
 *
 * Nsoraf Steering Of Roaming Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsoraf_SOR

// SteeringInfo - Contains either a PLMN-ID, an SNPN-ID or a GIN and, for the case of PLMNs, zero or more preferred access technologies for accessing such PLMN 
type SteeringInfo struct {

	PlmnId PlmnId `json:"plmnId,omitempty"`

	SnpnId PlmnIdNid `json:"snpnId,omitempty"`

	Gin PlmnIdNid `json:"gin,omitempty"`

	AccessTechList []AccessTech `json:"accessTechList,omitempty"`
}
