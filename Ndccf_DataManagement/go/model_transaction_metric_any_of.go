/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

type TransactionMetricAnyOf string

// List of TransactionMetricAnyOf
const (
	EST TransactionMetricAnyOf = "PDU_SES_EST"
	AUTH TransactionMetricAnyOf = "PDU_SES_AUTH"
	MODIF TransactionMetricAnyOf = "PDU_SES_MODIF"
	REL TransactionMetricAnyOf = "PDU_SES_REL"
)
