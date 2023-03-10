/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GenericNrm

type ManagedElementNcO struct {

	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`

	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`

	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`

	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`

	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`

	AlarmList AlarmListSingle `json:"AlarmList,omitempty"`

	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`

	Files []FilesSingle `json:"Files,omitempty"`
}
