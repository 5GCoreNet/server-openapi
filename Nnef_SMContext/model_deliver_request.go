/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnef_SMContext

import (
	"os"
)

type DeliverRequest struct {

	JsonData DeliverReqData `json:"jsonData,omitempty"`

	BinaryMoData *os.File `json:"binaryMoData,omitempty"`
}
