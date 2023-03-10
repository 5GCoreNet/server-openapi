/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsf_MBSUserDataIngestSession

// MbsDistSessionAnmt - Represents the set of MBS Distribution Session Announcements currently associated with this  MBS User Service Announcement. 
type MbsDistSessionAnmt struct {

	MbsSessionId MbsSessionId `json:"mbsSessionId,omitempty"`

	// MBS Frequency Selection Area Identifier
	MbsFSAId string `json:"mbsFSAId,omitempty"`

	DistrMethod DistributionMethod `json:"distrMethod"`

	ObjDistrAnnInfo ObjectDistMethAnmtInfo `json:"objDistrAnnInfo,omitempty"`

	SesDesInfo []string `json:"sesDesInfo"`
}
