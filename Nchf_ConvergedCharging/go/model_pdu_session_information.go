/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

import (
	"time"
)

type PduSessionInformation struct {

	NetworkSlicingInfo NetworkSlicingInfo `json:"networkSlicingInfo,omitempty"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionID int32 `json:"pduSessionID"`

	PduType PduSessionType `json:"pduType,omitempty"`

	SscMode SscMode `json:"sscMode,omitempty"`

	HPlmnId PlmnId `json:"hPlmnId,omitempty"`

	ServingNetworkFunctionID ServingNetworkFunctionId `json:"servingNetworkFunctionID,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	MAPDUNon3GPPRATType RatType `json:"mAPDUNon3GPPRATType,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	DnnId string `json:"dnnId"`

	DnnSelectionMode DnnSelectionMode `json:"dnnSelectionMode,omitempty"`

	ChargingCharacteristics string `json:"chargingCharacteristics,omitempty"`

	ChargingCharacteristicsSelectionMode ChargingCharacteristicsSelectionMode `json:"chargingCharacteristicsSelectionMode,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StopTime time.Time `json:"stopTime,omitempty"`

	Var3gppPSDataOffStatus Model3GpppsDataOffStatus `json:"3gppPSDataOffStatus,omitempty"`

	SessionStopIndicator bool `json:"sessionStopIndicator,omitempty"`

	PduAddress PduAddress `json:"pduAddress,omitempty"`

	Diagnostics int32 `json:"diagnostics,omitempty"`

	AuthorizedQoSInformation AuthorizedDefaultQos `json:"authorizedQoSInformation,omitempty"`

	SubscribedQoSInformation SubscribedDefaultQos `json:"subscribedQoSInformation,omitempty"`

	AuthorizedSessionAMBR Ambr `json:"authorizedSessionAMBR,omitempty"`

	SubscribedSessionAMBR Ambr `json:"subscribedSessionAMBR,omitempty"`

	ServingCNPlmnId PlmnId `json:"servingCNPlmnId,omitempty"`

	MAPDUSessionInformation MapduSessionInformation `json:"mAPDUSessionInformation,omitempty"`

	EnhancedDiagnostics []RanNasRelCause `json:"enhancedDiagnostics,omitempty"`

	RedundantTransmissionType RedundantTransmissionType `json:"redundantTransmissionType,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	PDUSessionPairID int32 `json:"pDUSessionPairID,omitempty"`

	CpCIoTOptimisationIndicator bool `json:"cpCIoTOptimisationIndicator,omitempty"`

	Var5GSControlPlaneOnlyIndicator bool `json:"5GSControlPlaneOnlyIndicator,omitempty"`

	SmallDataRateControlIndicator bool `json:"smallDataRateControlIndicator,omitempty"`

	Var5GLANTypeService Model5GlanTypeService `json:"5GLANTypeService,omitempty"`
}
