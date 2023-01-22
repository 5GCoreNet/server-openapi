/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AiMlNrm

type MlEntity struct {

	MLEntityId string `json:"mLEntityId,omitempty"`

	InferenceType string `json:"inferenceType,omitempty"`

	AIMLEntityVersion string `json:"aIMLEntityVersion,omitempty"`

	ExpectedRunTimeContext AimlContext `json:"expectedRunTimeContext,omitempty"`

	TrainingContext AimlContext `json:"trainingContext,omitempty"`

	RunTimeContext AimlContext `json:"runTimeContext,omitempty"`
}
