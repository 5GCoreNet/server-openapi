/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type TopologicalServiceArea struct {

	CellIdList []int32 `json:"cellIdList,omitempty"`

	TrackingAreaIdList []Tai1 `json:"trackingAreaIdList,omitempty"`

	ServingPLMN PlmnId1 `json:"servingPLMN,omitempty"`
}