/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

type NetworkPerfTypeAnyOf string

// List of NetworkPerfTypeAnyOf
const (
	GNB_ACTIVE_RATIO NetworkPerfTypeAnyOf = "GNB_ACTIVE_RATIO"
	GNB_COMPUTING_USAGE NetworkPerfTypeAnyOf = "GNB_COMPUTING_USAGE"
	GNB_MEMORY_USAGE NetworkPerfTypeAnyOf = "GNB_MEMORY_USAGE"
	GNB_DISK_USAGE NetworkPerfTypeAnyOf = "GNB_DISK_USAGE"
	NUM_OF_UE NetworkPerfTypeAnyOf = "NUM_OF_UE"
	SESS_SUCC_RATIO NetworkPerfTypeAnyOf = "SESS_SUCC_RATIO"
	HO_SUCC_RATIO NetworkPerfTypeAnyOf = "HO_SUCC_RATIO"
)
