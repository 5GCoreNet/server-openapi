/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFDiscovery

// CnfUnit - During the processing of cnfUnits attribute, all the members in the array shall be  interpreted as logically concatenated with logical \"AND\". 
type CnfUnit struct {

	CnfUnit []Atom `json:"cnfUnit"`
}
