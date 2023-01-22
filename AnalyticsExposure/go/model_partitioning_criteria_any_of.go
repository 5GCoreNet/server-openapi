/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

type PartitioningCriteriaAnyOf string

// List of PartitioningCriteriaAnyOf
const (
	TAC PartitioningCriteriaAnyOf = "TAC"
	SUBPLMN PartitioningCriteriaAnyOf = "SUBPLMN"
	GEOAREA PartitioningCriteriaAnyOf = "GEOAREA"
	SNSSAI PartitioningCriteriaAnyOf = "SNSSAI"
	DNN PartitioningCriteriaAnyOf = "DNN"
)