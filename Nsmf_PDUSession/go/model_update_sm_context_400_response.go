/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"os"
)

type UpdateSmContext400Response struct {

	JsonData SmContextUpdateError `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage *os.File `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmInformation *os.File `json:"binaryDataN2SmInformation,omitempty"`
}
