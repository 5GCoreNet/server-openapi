/*
 * TS 28.532 Streaming data reporting service
 *
 * OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package StreamingDataMnS

// TraceInfoType - Information specific to trace data reporting
type TraceInfoType struct {

	JobType JobTypeType `json:"jobType,omitempty"`

	ListOfInterfaces ListOfInterfacesType `json:"listOfInterfaces,omitempty"`

	// The Network Element types where Trace Session activation is needed. See 3GPP TS 32.422 clause 5.4 for additional details.
	ListOfNETypes []string `json:"listOfNETypes,omitempty"`

	PLMNTarget PLmnTargetType `json:"pLMNTarget,omitempty"`

	TraceReportingConsumerUri string `json:"traceReportingConsumerUri,omitempty"`

	TraceCollectionEntityIPAddress IpAddr `json:"traceCollectionEntityIPAddress,omitempty"`

	TraceDepth TraceDepthType `json:"traceDepth,omitempty"`

	TraceReference TraceReferenceType1 `json:"traceReference,omitempty"`

	TraceRecordingSessionReference string `json:"traceRecordingSessionReference,omitempty"`

	JobId string `json:"jobId,omitempty"`

	TraceReportingFormat TraceReportingFormatType `json:"traceReportingFormat,omitempty"`

	TraceTarget TraceTargetType `json:"traceTarget,omitempty"`

	TriggeringEvents TriggeringEventsType `json:"triggeringEvents,omitempty"`

	AnonymizationOfMDTData AnonymizationOfMdtDataType `json:"anonymizationOfMDTData,omitempty"`

	AreaConfigurationForNeighCell AreaConfig `json:"areaConfigurationForNeighCell,omitempty"`

	AreaScope []AreaScope `json:"areaScope,omitempty"`

	// Determines whether beam level measurements shall be included in case of immediate MDT M1 measurement in NR. For additional details see 3GPP TS 32.422 clause 5.10.40.
	BeamLevelMeasurement bool `json:"beamLevelMeasurement,omitempty"`

	CollectionPeriodRRMLTE CollectionPeriodRrmlteType `json:"collectionPeriodRRMLTE,omitempty"`

	CollectionPeriodM6LTE CollectionPeriodM6LteType `json:"collectionPeriodM6LTE,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.33.
	CollectionPeriodM7LTE int32 `json:"collectionPeriodM7LTE,omitempty"`

	CollectionPeriodRRMUMTS CollectionPeriodRrmumtsType `json:"collectionPeriodRRMUMTS,omitempty"`

	CollectionPeriodRRMNR CollectionPeriodRrmnrType `json:"collectionPeriodRRMNR,omitempty"`

	CollectionPeriodM6NR CollectionPeriodM6NrType `json:"collectionPeriodM6NR,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.35.
	CollectionPeriodM7NR int32 `json:"collectionPeriodM7NR,omitempty"`

	EventListForEventTriggeredMeasurement EventListForEventTriggeredMeasurementType `json:"eventListForEventTriggeredMeasurement,omitempty"`

	EventThreshold EventThresholdType `json:"eventThreshold,omitempty"`

	ListOfMeasurements ListOfMeasurementsType `json:"listOfMeasurements,omitempty"`

	LoggingDuration LoggingDurationType `json:"loggingDuration,omitempty"`

	LoggingInterval LoggingIntervalType `json:"loggingInterval,omitempty"`

	EventThresholdL1 EventThresholdL1Type `json:"eventThresholdL1,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.Y.
	HysteresisL1 int32 `json:"hysteresisL1,omitempty"`

	TimeToTriggerL1 TimeToTriggerL1Type `json:"timeToTriggerL1,omitempty"`

	MBSFNAreaList []MbsfnArea `json:"mBSFNAreaList,omitempty"`

	MeasurementPeriodLTE MeasurementPeriodLteType `json:"measurementPeriodLTE,omitempty"`

	MeasurementPeriodUMTS MeasurementPeriodUmtsType `json:"measurementPeriodUMTS,omitempty"`

	MeasurementQuantity MeasurementQuantityType `json:"measurementQuantity,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.A.
	EventThresholdUphUMTS int32 `json:"eventThresholdUphUMTS,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.24.
	PLMNList []PLmnListTypeInner `json:"pLMNList,omitempty"`

	PositioningMethod PositioningMethodType `json:"positioningMethod,omitempty"`

	ReportAmount ReportAmountType `json:"reportAmount,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.4.
	ReportingTrigger []string `json:"reportingTrigger,omitempty"`

	ReportInterval ReportIntervalType `json:"reportInterval,omitempty"`

	ReportType ReportTypeType `json:"reportType,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.29.
	SensorInformation []string `json:"sensorInformation,omitempty"`

	// See details in 3GPP TS 32.422 clause 5.10.11. Only TCE Id value may be sent over the air to the UE being configured for Logged MDT.
	TraceCollectionEntityId int32 `json:"traceCollectionEntityId,omitempty"`
}
