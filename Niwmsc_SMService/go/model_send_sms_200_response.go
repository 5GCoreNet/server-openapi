/*
 * Niwmsc_SMService
 *
 * SMS-IWMSC Short Message Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Niwmsc_SMService

import (
	"os"
)

type SendSms200Response struct {

	JsonData SmsDeliveryData `json:"jsonData,omitempty"`

	BinaryPayload *os.File `json:"binaryPayload,omitempty"`
}