/*
 * Intent NRM
 *
 * OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package IntentNrm

// WeakRsrpContext - This data type is the \"TargetContext\" data type with specialisations for WeakRSRPContext
type WeakRsrpContext struct {

	ContextAttribute string `json:"contextAttribute,omitempty"`

	ContextCondition string `json:"contextCondition,omitempty"`

	ContextValueRange float32 `json:"contextValueRange,omitempty"`
}
