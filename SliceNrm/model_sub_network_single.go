/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

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

	NetworkSlice []NetworkSliceSingle `json:"NetworkSlice,omitempty"`

	NetworkSliceSubnet []NetworkSliceSubnetSingle `json:"NetworkSliceSubnet,omitempty"`

	EPTransport []EpTransportSingle `json:"EP_Transport,omitempty"`

	NetworkSliceSubnetProviderCapabilities []NetworkSliceSubnetProviderCapabilitiesSingle `json:"NetworkSliceSubnetProviderCapabilities,omitempty"`

	FeasibilityCheckAndReservationJob []FeasibilityCheckAndReservationJobSingle `json:"FeasibilityCheckAndReservationJob,omitempty"`
}
