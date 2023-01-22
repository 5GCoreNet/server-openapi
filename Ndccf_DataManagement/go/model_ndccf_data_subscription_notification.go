/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

import (
	"time"
)

// NdccfDataSubscriptionNotification - Represents a notification for a DCCF data subscription.
type NdccfDataSubscriptionNotification struct {

	// Notification correlation identifier.
	DataNotifCorrId string `json:"dataNotifCorrId"`

	DataNotif DataNotification `json:"dataNotif,omitempty"`

	// List of reports with summarized data from multiple notifications received from data producer. 
	DataReports []NotifSummaryReport `json:"dataReports,omitempty"`

	FetchInstruct FetchInstruction `json:"fetchInstruct,omitempty"`

	// It indicates the termination of the data management subscription that requested by the DCCF. 
	TerminationReq bool `json:"terminationReq,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}
