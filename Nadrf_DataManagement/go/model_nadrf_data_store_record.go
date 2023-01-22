/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// NadrfDataStoreRecord - Represents an Individual ADRF Data Store Record.
type NadrfDataStoreRecord struct {

	DataNotif DataNotification `json:"dataNotif,omitempty"`

	// List of analytics subscription notifications.
	AnaNotifications []NnwdafEventsSubscriptionNotification `json:"anaNotifications,omitempty"`

	// Represents the subscription information of the corresponding analytics notification. 
	AnaSub []NnwdafEventsSubscription `json:"anaSub,omitempty"`

	// Represents the subscription information of the corresponding data notification. 
	DataSub []DataSubscription `json:"dataSub,omitempty"`
}