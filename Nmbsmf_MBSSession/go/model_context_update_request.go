/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbsmf_MBSSession

import (
	"os"
)

type ContextUpdateRequest struct {

	JsonData ContextUpdateReqData `json:"jsonData,omitempty"`

	BinaryDataN2Information *os.File `json:"binaryDataN2Information,omitempty"`
}
