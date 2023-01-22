/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// DataNotification - Represents a Data Subscription Notification.
type DataNotification struct {

	// List of notifications of AMF events.
	AmfEventNotifs []AmfEventNotification `json:"amfEventNotifs,omitempty"`

	// List of notifications of SMF events.
	SmfEventNotifs []NsmfEventExposureNotification `json:"smfEventNotifs,omitempty"`

	// List of notifications of UDM events.
	UdmEventNotifs []MonitoringReport `json:"udmEventNotifs,omitempty"`

	// List of notifications of NEF events.
	NefEventNotifs []NefEventExposureNotif `json:"nefEventNotifs,omitempty"`

	// List of notifications of AF events.
	AfEventNotifs []AfEventExposureNotif `json:"afEventNotifs,omitempty"`

	// List of notifications of NRF events.
	NrfEventNotifs []NotificationData `json:"nrfEventNotifs,omitempty"`

	// List of notifications of NSACF events.
	NsacfEventNotifs []SacEventReport `json:"nsacfEventNotifs,omitempty"`
}
