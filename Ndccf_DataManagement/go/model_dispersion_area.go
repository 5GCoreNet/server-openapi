/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// DispersionArea - Dispersion Area
type DispersionArea struct {

	TaiList []Tai `json:"taiList,omitempty"`

	NcgiList []Ncgi `json:"ncgiList,omitempty"`

	EcgiList []Ecgi `json:"ecgiList,omitempty"`

	N3gaInd bool `json:"n3gaInd,omitempty"`
}
