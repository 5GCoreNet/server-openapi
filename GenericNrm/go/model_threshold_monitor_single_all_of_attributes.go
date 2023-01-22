/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package GenericNrm

type ThresholdMonitorSingleAllOfAttributes struct {

	AdministrativeState AdministrativeState `json:"administrativeState,omitempty"`

	OperationalState OperationalState `json:"operationalState,omitempty"`

	PerformanceMetrics []string `json:"performanceMetrics,omitempty"`

	ThresholdInfoList []ThresholdInfo `json:"thresholdInfoList,omitempty"`

	MonitorGranularityPeriod int32 `json:"monitorGranularityPeriod,omitempty"`

	ObjectInstances []string `json:"objectInstances,omitempty"`

	RootObjectInstances []string `json:"rootObjectInstances,omitempty"`
}
