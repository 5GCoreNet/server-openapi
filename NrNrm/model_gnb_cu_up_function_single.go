/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type GnbCuUpFunctionSingle struct {

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes ManagedFunctionAttr `json:"attributes,omitempty"`

	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`

	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`

	ManagedNFService []ManagedNfServiceSingle `json:"ManagedNFService,omitempty"`

	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`

	RRMPolicyRatio []RrmPolicyRatioSingle `json:"RRMPolicyRatio,omitempty"`

	EPE1 EpE1Single `json:"EP_E1,omitempty"`

	EPXnU []EpXnUSingle `json:"EP_XnU,omitempty"`

	EPF1U []EpF1USingle `json:"EP_F1U,omitempty"`

	EPNgU []EpNgUSingle `json:"EP_NgU,omitempty"`

	EPX2U []EpX2USingle `json:"EP_X2U,omitempty"`

	EPS1U []EpS1USingle `json:"EP_S1U,omitempty"`
}