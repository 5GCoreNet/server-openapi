/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_UELocation

type PartitioningCriteriaAnyOf string

// List of PartitioningCriteriaAnyOf
const (
	TAC PartitioningCriteriaAnyOf = "TAC"
	SUBPLMN PartitioningCriteriaAnyOf = "SUBPLMN"
	GEOAREA PartitioningCriteriaAnyOf = "GEOAREA"
	SNSSAI PartitioningCriteriaAnyOf = "SNSSAI"
	DNN PartitioningCriteriaAnyOf = "DNN"
)
