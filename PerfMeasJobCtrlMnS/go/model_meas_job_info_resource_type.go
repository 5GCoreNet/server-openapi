/*
 * TS 28.550 Performance Measurement Job Control Service
 *
 * OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 16.8.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package PerfMeasJobCtrlMnS

type MeasJobInfoResourceType struct {

	Href string `json:"href,omitempty"`

	IOCName string `json:"iOCName,omitempty"`

	IOCInstanceList []string `json:"iOCInstanceList,omitempty"`

	MeasurementCategoryList []string `json:"measurementCategoryList,omitempty"`

	ReportingMethod ReportingMethodType `json:"reportingMethod,omitempty"`

	GranularityPeriod int32 `json:"granularityPeriod,omitempty"`

	ReportingPeriod int32 `json:"reportingPeriod,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	StopTime string `json:"stopTime,omitempty"`

	Schedule ScheduleType `json:"schedule,omitempty"`

	StreamTarget string `json:"streamTarget,omitempty"`

	Priority PriorityType `json:"priority,omitempty"`

	Reliability string `json:"reliability,omitempty"`
}
