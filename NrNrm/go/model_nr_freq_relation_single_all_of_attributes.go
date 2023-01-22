/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NrNrm

type NrFreqRelationSingleAllOfAttributes struct {

	OffsetMO QOffsetRangeList `json:"offsetMO,omitempty"`

	BlockListEntry []int32 `json:"blockListEntry,omitempty"`

	BlockListEntryIdleMode int32 `json:"blockListEntryIdleMode,omitempty"`

	CellReselectionPriority int32 `json:"cellReselectionPriority,omitempty"`

	CellReselectionSubPriority float32 `json:"cellReselectionSubPriority,omitempty"`

	PMax int32 `json:"pMax,omitempty"`

	QOffsetFreq float32 `json:"qOffsetFreq,omitempty"`

	QQualMin float32 `json:"qQualMin,omitempty"`

	QRxLevMin int32 `json:"qRxLevMin,omitempty"`

	ThreshXHighP int32 `json:"threshXHighP,omitempty"`

	ThreshXHighQ int32 `json:"threshXHighQ,omitempty"`

	ThreshXLowP int32 `json:"threshXLowP,omitempty"`

	ThreshXLowQ int32 `json:"threshXLowQ,omitempty"`

	TReselectionNr int32 `json:"tReselectionNr,omitempty"`

	TReselectionNRSfHigh TReselectionNrsf `json:"tReselectionNRSfHigh,omitempty"`

	TReselectionNRSfMedium TReselectionNrsf `json:"tReselectionNRSfMedium,omitempty"`

	NRFrequencyRef string `json:"nRFrequencyRef,omitempty"`
}
