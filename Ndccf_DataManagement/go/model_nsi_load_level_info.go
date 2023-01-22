/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// NsiLoadLevelInfo - Represents the network slice and optionally the associated network slice instance and the  load level information. 
type NsiLoadLevelInfo struct {

	// Load level information of the network slice and the optionally associated network slice  instance. 
	LoadLevelInformation int32 `json:"loadLevelInformation"`

	Snssai Snssai `json:"snssai"`

	// Contains the Identifier of the selected Network Slice instance
	NsiId string `json:"nsiId,omitempty"`

	ResUsage ResourceUsage `json:"resUsage,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumOfExceedLoadLevelThr int32 `json:"numOfExceedLoadLevelThr,omitempty"`

	ExceedLoadLevelThrInd bool `json:"exceedLoadLevelThrInd,omitempty"`

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	TimePeriod TimeWindow `json:"timePeriod,omitempty"`

	// Each element indicates the time elapsed between times each threshold is met or exceeded or crossed. 
	ResUsgThrCrossTimePeriod []TimeWindow `json:"resUsgThrCrossTimePeriod,omitempty"`

	NumOfUes NumberAverage `json:"numOfUes,omitempty"`

	NumOfPduSess NumberAverage `json:"numOfPduSess,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}