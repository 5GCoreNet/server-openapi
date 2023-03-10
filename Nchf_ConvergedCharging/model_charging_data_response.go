/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

import (
	"time"
)

type ChargingDataResponse struct {

	// string with format 'date-time' as defined in OpenAPI.
	InvocationTimeStamp time.Time `json:"invocationTimeStamp"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	InvocationSequenceNumber int32 `json:"invocationSequenceNumber"`

	InvocationResult InvocationResult `json:"invocationResult,omitempty"`

	SessionFailover SessionFailover `json:"sessionFailover,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	MultipleUnitInformation []MultipleUnitInformation `json:"multipleUnitInformation,omitempty"`

	Triggers []Trigger `json:"triggers,omitempty"`

	PDUSessionChargingInformation PduSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty"`

	RoamingQBCInformation RoamingQbcInformation `json:"roamingQBCInformation,omitempty"`

	LocationReportingChargingInformation LocationReportingChargingInformation `json:"locationReportingChargingInformation,omitempty"`
}
