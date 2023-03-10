/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type ManagedElementSingle struct {

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes ManagedElementAttr `json:"attributes,omitempty"`

	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`

	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`

	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`

	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`

	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`

	AlarmList AlarmListSingle `json:"AlarmList,omitempty"`

	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`

	Files []FilesSingle `json:"Files,omitempty"`

	GnbDuFunction []GnbDuFunctionSingle `json:"GnbDuFunction,omitempty"`

	GnbCuUpFunction []GnbCuUpFunctionSingle `json:"GnbCuUpFunction,omitempty"`

	GnbCuCpFunction []GnbCuCpFunctionSingle `json:"GnbCuCpFunction,omitempty"`

	DESManagementFunction DesManagementFunctionSingle `json:"DESManagementFunction,omitempty"`

	DRACHOptimizationFunction DrachOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`

	DMROFunction DmroFunctionSingle `json:"DMROFunction,omitempty"`

	DLBOFunction DlboFunctionSingle `json:"DLBOFunction,omitempty"`

	DPCIConfigurationFunction DpciConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`

	CPCIConfigurationFunction CpciConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`

	CESManagementFunction CesManagementFunctionSingle `json:"CESManagementFunction,omitempty"`

	Configurable5QISet []Configurable5QiSetSingle `json:"Configurable5QISet,omitempty"`

	Dynamic5QISet []Dynamic5QiSetSingle `json:"Dynamic5QISet,omitempty"`
}
