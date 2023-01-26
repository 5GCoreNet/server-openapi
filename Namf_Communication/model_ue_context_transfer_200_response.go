/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

import (
	"os"
)

type UeContextTransfer200Response struct {

	JsonData UeContextTransferRspData `json:"jsonData,omitempty"`

	BinaryDataN2Information *os.File `json:"binaryDataN2Information,omitempty"`

	BinaryDataN2InformationExt1 *os.File `json:"binaryDataN2InformationExt1,omitempty"`

	BinaryDataN2InformationExt2 *os.File `json:"binaryDataN2InformationExt2,omitempty"`
}