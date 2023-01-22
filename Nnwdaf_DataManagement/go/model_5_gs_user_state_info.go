/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// Model5GsUserStateInfo - Represents the 5GS User state of the UE for an access type
type Model5GsUserStateInfo struct {

	Var5gsUserState Model5GsUserState `json:"5gsUserState"`

	AccessType AccessType `json:"accessType"`
}
