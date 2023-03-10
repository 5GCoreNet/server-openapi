/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

type CreateMonitoringEventSubscription200Response struct {

	ImeiChange AssociationType `json:"imeiChange,omitempty"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	IdleStatusInfo IdleStatusInfo `json:"idleStatusInfo,omitempty"`

	LocationInfo LocationInfo `json:"locationInfo,omitempty"`

	LocFailureCause LocationFailureCause `json:"locFailureCause,omitempty"`

	// If \"monitoringType\" is \"LOSS_OF_CONNECTIVITY\", this parameter shall be included if available to identify the reason why loss of connectivity is reported. Refer to 3GPP TS 29.336 clause 8.4.58.
	LossOfConnectReason int32 `json:"lossOfConnectReason,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxUEAvailabilityTime time.Time `json:"maxUEAvailabilityTime,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	MonitoringType MonitoringType `json:"monitoringType"`

	UePerLocationReport UePerLocationReport `json:"uePerLocationReport,omitempty"`

	PlmnId PlmnId `json:"plmnId,omitempty"`

	ReachabilityType ReachabilityType `json:"reachabilityType,omitempty"`

	// If \"monitoringType\" is \"ROAMING_STATUS\", this parameter shall be set to \"true\" if the UE is on roaming status. Set to false or omitted otherwise.
	RoamingStatus bool `json:"roamingStatus,omitempty"`

	FailureCause FailureCause `json:"failureCause,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	EventTime time.Time `json:"eventTime,omitempty"`

	PdnConnInfoList []PdnConnectionInformation `json:"pdnConnInfoList,omitempty"`

	DddStatus DlDataDeliveryStatus `json:"dddStatus,omitempty"`

	DddTrafDescriptor DddTrafficDescriptor `json:"dddTrafDescriptor,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	MaxWaitTime time.Time `json:"maxWaitTime,omitempty"`

	ApiCaps []ApiCapabilityInfo `json:"apiCaps,omitempty"`

	NSStatusInfo SacEventStatus `json:"nSStatusInfo,omitempty"`

	AfServiceId string `json:"afServiceId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter may be included to identify the UAV.
	ServLevelDevId string `json:"servLevelDevId,omitempty"`

	// If \"monitoringType\" is \"AREA_OF_INTEREST\", this parameter shall be set to true if the specified UAV is in the monitoring area. Set to false or omitted otherwise.
	UavPresInd bool `json:"uavPresInd,omitempty"`

	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports"`
}
