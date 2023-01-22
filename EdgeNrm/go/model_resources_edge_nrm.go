/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EdgeNrm

type ResourcesEdgeNrm struct {

	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes ManagedFunctionAttr `json:"attributes,omitempty"`

	Subnetwork []SubNetworkSingle `json:"Subnetwork,omitempty"`

	ECSFunction []EcsFunctionSingle `json:"ECSFunction,omitempty"`

	EdgeDataNetwork []EdgeDataNetworkSingle `json:"EdgeDataNetwork,omitempty"`

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

	ManagedNFService []ManagedNfServiceSingle `json:"ManagedNFService,omitempty"`

	EdnIdentifier string `json:"ednIdentifier,omitempty"`

	EDNConnectionInfo EdnConnectionInfo `json:"eDNConnectionInfo,omitempty"`

	EASFunction []EasFunctionSingle `json:"EASFunction,omitempty"`

	EESFunction []EesFunctionSingle `json:"EESFunction,omitempty"`

	RequiredEASservingLocation ServingLocation `json:"requiredEASservingLocation,omitempty"`

	AffinityAntiAffinity AffinityAntiAffinity `json:"affinityAntiAffinity,omitempty"`

	ServiceContinuity bool `json:"serviceContinuity,omitempty"`

	VirtualResource VirtualResource `json:"virtualResource,omitempty"`

	SoftwareImageInfo SoftwareImageInfo `json:"softwareImageInfo,omitempty"`
}
