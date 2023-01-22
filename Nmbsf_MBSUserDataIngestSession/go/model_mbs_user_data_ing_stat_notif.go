/*
 * nmbsf-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsf_MBSUserDataIngestSession

// MbsUserDataIngStatNotif - Represents an MBS User Data Ingest Session Status Notification. 
type MbsUserDataIngStatNotif struct {

	MbsIngSessionId string `json:"mbsIngSessionId"`

	EventNotifs []EventNotification `json:"eventNotifs"`
}
