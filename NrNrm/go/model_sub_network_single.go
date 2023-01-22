/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NrNrm

type SubNetworkSingle struct {

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes SubNetworkAttr `json:"attributes,omitempty"`

	ManagementNode []ManagementNodeSingle `json:"ManagementNode,omitempty"`

	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`

	MeContext []MeContextSingle `json:"MeContext,omitempty"`

	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`

	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`

	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`

	ManagementDataCollection []ManagementDataCollectionSingle `json:"ManagementDataCollection,omitempty"`

	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`

	AlarmList AlarmListSingle `json:"AlarmList,omitempty"`

	Files []FilesSingle `json:"Files,omitempty"`

	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`

	MnsRegistry MnsRegistrySingle `json:"MnsRegistry,omitempty"`

	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`

	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`

	NRFrequency []NrFrequencySingle `json:"NRFrequency,omitempty"`

	ExternalGnbCuCpFunction []ExternalGnbCuCpFunctionSingle `json:"ExternalGnbCuCpFunction,omitempty"`

	ExternalENBFunction []ExternalEnbFunctionSingle `json:"ExternalENBFunction,omitempty"`

	EUtranFrequency []EUtranFrequencySingle `json:"EUtranFrequency,omitempty"`

	DESManagementFunction DesManagementFunctionSingle `json:"DESManagementFunction,omitempty"`

	DRACHOptimizationFunction DrachOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`

	DMROFunction DmroFunctionSingle `json:"DMROFunction,omitempty"`

	DLBOFunction DlboFunctionSingle `json:"DLBOFunction,omitempty"`

	DPCIConfigurationFunction DpciConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`

	CPCIConfigurationFunction CpciConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`

	CESManagementFunction CesManagementFunctionSingle `json:"CESManagementFunction,omitempty"`

	Configurable5QISet []Configurable5QiSetSingle `json:"Configurable5QISet,omitempty"`

	RimRSGlobal RimRsGlobalSingle `json:"RimRSGlobal,omitempty"`

	Dynamic5QISet []Dynamic5QiSetSingle `json:"Dynamic5QISet,omitempty"`

	CCOFunction CcoFunctionSingle `json:"CCOFunction,omitempty"`
}
