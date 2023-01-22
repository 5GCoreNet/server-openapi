/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsf_MBSUserDataIngestSession

// MbsUserDataIngSessionPatch - Represents the requested modifications to an MBS User Data Ingest Session Status  Subscription. 
type MbsUserDataIngSessionPatch struct {

	// Contains the requested modifications to one or more MBS Distribution Session(s)  composing the MBS User Data Ingest Session. 
	MbsDistSessInfos map[string]MbsDistributionSessionInfo `json:"mbsDistSessInfos,omitempty"`

	ActPeriods []TimeWindow `json:"actPeriods,omitempty"`
}
