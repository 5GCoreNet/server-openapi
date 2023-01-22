/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_Events

// LocationDevMonReport - Location deviation monitoring report.
type LocationDevMonReport struct {

	// List of VAL Users or UE IDs for which report is related to.
	TgtUes []ValTargetUe `json:"tgtUes"`

	LocInfo LocationInfo `json:"locInfo"`

	NotifType LocDevNotification `json:"notifType"`
}