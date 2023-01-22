/*
 * Nhss_UECM
 *
 * HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_UECM

// RoamingStatusUpdateInfo - It represents the request body of the Roaming Status Update request sent by UDM to HSS, and contains the IMSI of the UE and the new PLMN-ID 
type RoamingStatusUpdateInfo struct {

	Imsi string `json:"imsi"`

	PlmnId PlmnId `json:"plmnId"`
}
