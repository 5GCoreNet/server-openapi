/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

import (
	"time"
)

// NdccfAnalyticsSubscriptionNotification - Represents a notification for a DCCF analytics subscription.
type NdccfAnalyticsSubscriptionNotification struct {

	// Notification correlation identifier.
	AnaNotifCorrId string `json:"anaNotifCorrId"`

	// List of analytics subscription notifications.
	AnaNotifications []NnwdafEventsSubscriptionNotification `json:"anaNotifications,omitempty"`

	// List of reports with summarized data from multiple analytics notifications that the DCCF has received from NWDAF. 
	AnaReports []NotifSummaryReport `json:"anaReports,omitempty"`

	FetchInstruct FetchInstruction `json:"fetchInstruct,omitempty"`

	// It indicates the termination of the data management subscription that requested by the DCCF. 
	TerminationReq bool `json:"terminationReq,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}
