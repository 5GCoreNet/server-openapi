/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// NadrfDataRetrievalSubscription - Represents an Individual ADRF Data Retrieval Subscription.
type NadrfDataRetrievalSubscription struct {

	AnaSub NnwdafEventsSubscription `json:"anaSub,omitempty"`

	DataSub DataSubscription `json:"dataSub,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotificationURI string `json:"notificationURI"`

	TimePeriod TimeWindow `json:"timePeriod"`

	// Notification correlation identifier.
	NotifCorrId string `json:"notifCorrId"`
}
