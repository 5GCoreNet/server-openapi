/*
 * ECS EES Registration_API
 *
 * API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eecs_EESRegistration

type AcrScenarioAnyOf string

// List of AcrScenarioAnyOf
const (
	EEC_INITIATED AcrScenarioAnyOf = "EEC_INITIATED"
	EEC_EXECUTED_VIA_SOURCE_EES AcrScenarioAnyOf = "EEC_EXECUTED_VIA_SOURCE_EES"
	EEC_EXECUTED_VIA_TARGET_EES AcrScenarioAnyOf = "EEC_EXECUTED_VIA_TARGET_EES"
	SOURCE_EAS_DECIDED AcrScenarioAnyOf = "SOURCE_EAS_DECIDED"
	SOURCE_EES_EXECUTED AcrScenarioAnyOf = "SOURCE_EES_EXECUTED"
	EEL_MANAGED_ACR AcrScenarioAnyOf = "EEL_MANAGED_ACR"
)
