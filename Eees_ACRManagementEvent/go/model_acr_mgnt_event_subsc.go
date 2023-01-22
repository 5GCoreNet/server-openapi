/*
 * EES ACR Management Event_API
 *
 * API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_ACRManagementEvent

// AcrMgntEventSubsc - Represents an ACR Management Event Subscription.
type AcrMgntEventSubsc struct {

	Event AcrMgntEvent `json:"event"`

	EventFilter AcrMgntEventFilter `json:"eventFilter,omitempty"`

	EvtReq ReportingInformation `json:"evtReq,omitempty"`

	TgtUeId TargetUeIdentification `json:"tgtUeId,omitempty"`

	DnaiChgType DnaiChangeType `json:"dnaiChgType,omitempty"`

	EasAckInd bool `json:"easAckInd,omitempty"`

	// A list of EAS characteristics.
	EasChars []EasCharacteristics `json:"easChars,omitempty"`
}
