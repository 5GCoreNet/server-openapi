/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

type PartitioningCriteriaAnyOf string

// List of PartitioningCriteriaAnyOf
const (
	TAC PartitioningCriteriaAnyOf = "TAC"
	SUBPLMN PartitioningCriteriaAnyOf = "SUBPLMN"
	GEOAREA PartitioningCriteriaAnyOf = "GEOAREA"
	SNSSAI PartitioningCriteriaAnyOf = "SNSSAI"
	DNN PartitioningCriteriaAnyOf = "DNN"
)
