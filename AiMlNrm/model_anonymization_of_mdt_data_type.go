/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AiMlNrm
// AnonymizationOfMdtDataType : Specifies level of MDT anonymization. For additional details see 3GPP TS 32.422 clause 5.10.12.
type AnonymizationOfMdtDataType string

// List of AnonymizationOfMdtDataType
const (
	NO_IDENTITY AnonymizationOfMdtDataType = "NO_IDENTITY"
	TAC_OF_IMEI AnonymizationOfMdtDataType = "TAC_OF_IMEI"
)
