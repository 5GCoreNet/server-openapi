/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package GenericNrm

type AreaConfig struct {

	FreqInfo FreqInfo `json:"freqInfo,omitempty"`

	PciList []int32 `json:"pciList,omitempty"`
}