/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// PlmnOauth2 - Oauth2.0 required indication for a given PLMN ID
type PlmnOauth2 struct {

	Oauth2RequiredPlmnIdList []PlmnId `json:"oauth2RequiredPlmnIdList,omitempty"`

	Oauth2NotRequiredPlmnIdList []PlmnId `json:"oauth2NotRequiredPlmnIdList,omitempty"`
}
