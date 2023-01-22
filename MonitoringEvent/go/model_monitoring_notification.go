/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MonitoringEvent

// MonitoringNotification - Represents an event monitoring notification.
type MonitoringNotification struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Subscription string `json:"subscription"`

	// Each element identifies a notification of grouping configuration result.
	ConfigResults []ConfigResult `json:"configResults,omitempty"`

	// Monitoring event reports.
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports,omitempty"`

	// Identifies the added external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedExternalIds []string `json:"addedExternalIds,omitempty"`

	// Identifies the added MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	AddedMsisdns []string `json:"addedMsisdns,omitempty"`

	// Identifies the cancelled external Identifier(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelExternalIds []string `json:"cancelExternalIds,omitempty"`

	// Identifies the cancelled MSISDN(s) within the active group via the \"externalGroupId\" attribute within the MonitoringEventSubscription data type.
	CancelMsisdns []string `json:"cancelMsisdns,omitempty"`

	// Indicates whether to request to cancel the corresponding monitoring subscription. Set to false or omitted otherwise. 
	CancelInd bool `json:"cancelInd,omitempty"`

	AppliedParam AppliedParameterConfiguration `json:"appliedParam,omitempty"`
}
