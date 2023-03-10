/*
 * TS 28.532 Performance Threshold Monitoring MnS
 *
 * OAS 3.0.1 definition of the Performance Threshold Monitoring MnS © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_PerfMnS

type NotifyThresholdCrossing struct {

	Href string `json:"href"`

	NotificationId int32 `json:"notificationId"`

	NotificationType NotificationType `json:"notificationType"`

	EventTime time.Time `json:"eventTime"`

	SystemDN string `json:"systemDN"`

	ObservedPerfMetricName string `json:"observedPerfMetricName,omitempty"`

	ObservedPerfMetricValue PerfMetricValue `json:"observedPerfMetricValue,omitempty"`

	ObservedPerfMetricDirection PerfMetricDirection `json:"observedPerfMetricDirection,omitempty"`

	ThresholdValue PerfMetricValue `json:"thresholdValue,omitempty"`

	Hysteresis PerfMetricValue `json:"hysteresis,omitempty"`

	MonitorGranularityPeriod int32 `json:"monitorGranularityPeriod,omitempty"`

	AdditionalText string `json:"additionalText,omitempty"`
}
