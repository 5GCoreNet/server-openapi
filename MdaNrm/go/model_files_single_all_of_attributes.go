/*
 * MDA NRM
 *
 * OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MdaNrm

type FilesSingleAllOfAttributes struct {

	NumberOfFiles int32 `json:"numberOfFiles,omitempty"`

	JobRef string `json:"jobRef,omitempty"`

	JobId string `json:"jobId,omitempty"`

	File []FileSingle `json:"File,omitempty"`
}
