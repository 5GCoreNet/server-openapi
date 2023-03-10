/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AiMlNrm

type SubNetworkAttr struct {

	DnPrefix string `json:"dnPrefix,omitempty"`

	UserLabel string `json:"userLabel,omitempty"`

	UserDefinedNetworkType string `json:"userDefinedNetworkType,omitempty"`

	SetOfMcc []string `json:"setOfMcc,omitempty"`

	PriorityLabel int32 `json:"priorityLabel,omitempty"`

	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`

	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}
