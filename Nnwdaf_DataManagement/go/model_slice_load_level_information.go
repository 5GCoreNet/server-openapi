/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// SliceLoadLevelInformation - Contains load level information applicable for one or several slices.
type SliceLoadLevelInformation struct {

	// Load level information of the network slice and the optionally associated network slice  instance. 
	LoadLevelInformation int32 `json:"loadLevelInformation"`

	// Identification(s) of network slice to which the subscription applies.
	Snssais []Snssai `json:"snssais"`
}
