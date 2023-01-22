/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

type CreditManagementStatusAnyOf string

// List of CreditManagementStatusAnyOf
const (
	END_USER_SER_DENIED CreditManagementStatusAnyOf = "END_USER_SER_DENIED"
	CREDIT_CTRL_NOT_APP CreditManagementStatusAnyOf = "CREDIT_CTRL_NOT_APP"
	AUTH_REJECTED CreditManagementStatusAnyOf = "AUTH_REJECTED"
	USER_UNKNOWN CreditManagementStatusAnyOf = "USER_UNKNOWN"
	RATING_FAILED CreditManagementStatusAnyOf = "RATING_FAILED"
)
