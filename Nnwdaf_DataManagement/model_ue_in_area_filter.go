/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// UeInAreaFilter - Additional filters for UE in Area Report event
type UeInAreaFilter struct {

	UeType UeType `json:"ueType,omitempty"`

	AerialSrvDnnInd bool `json:"aerialSrvDnnInd,omitempty"`
}
