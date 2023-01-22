/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_UECM

import (
	"time"
)

// DataRestorationNotification - Contains identities representing those UEs potentially affected by a data-loss event at the UDR
type DataRestorationNotification struct {

	// string with format 'date-time' as defined in OpenAPI.
	LastReplicationTime time.Time `json:"lastReplicationTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ResetIds []string `json:"resetIds,omitempty"`

	SNssaiList []Snssai `json:"sNssaiList,omitempty"`

	DnnList []string `json:"dnnList,omitempty"`

	// Identifier of a group of NFs.
	UdmGroupId string `json:"udmGroupId,omitempty"`
}
