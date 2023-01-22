/*
 * coslaNrm
 *
 * OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CoslaNrm

type AssuranceGoalStatus struct {

	AssuranceGoalStatusId string `json:"assuranceGoalStatusId,omitempty"`

	AssuranceGoalId string `json:"assuranceGoalId,omitempty"`

	AssuranceGoalStatusObserved AssuranceGoalStatusObserved `json:"assuranceGoalStatusObserved,omitempty"`

	AssuranceGoalStatusPredicted AssuranceGoalStatusPredicted `json:"assuranceGoalStatusPredicted,omitempty"`

	AssuranceGoalRef string `json:"assuranceGoalRef,omitempty"`
}
