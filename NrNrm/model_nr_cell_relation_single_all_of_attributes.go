/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_NrNrm

type NrCellRelationSingleAllOfAttributes struct {

	NRTCI int32 `json:"nRTCI,omitempty"`

	CellIndividualOffset CellIndividualOffset `json:"cellIndividualOffset,omitempty"`

	AdjacentNRCellRef string `json:"adjacentNRCellRef,omitempty"`

	NRFreqRelationRef string `json:"nRFreqRelationRef,omitempty"`

	IsRemoveAllowed bool `json:"isRemoveAllowed,omitempty"`

	IsHOAllowed bool `json:"isHOAllowed,omitempty"`

	IsESCoveredBy IsEsCoveredBy `json:"isESCoveredBy,omitempty"`

	IsENDCAllowed bool `json:"isENDCAllowed,omitempty"`

	IsMLBAllowed bool `json:"isMLBAllowed,omitempty"`
}