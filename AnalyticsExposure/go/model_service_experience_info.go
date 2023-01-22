/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

// ServiceExperienceInfo - Represents service experience information.
type ServiceExperienceInfo struct {

	SvcExprc SvcExperience `json:"svcExprc"`

	// string with format 'float' as defined in OpenAPI.
	SvcExprcVariance float32 `json:"svcExprcVariance,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	SrvExpcType ServiceExperienceType `json:"srvExpcType,omitempty"`

	UeLocs []LocationInfo `json:"ueLocs,omitempty"`

	UpfInfo UpfInformation `json:"upfInfo,omitempty"`

	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai,omitempty"`

	AppServerInst AddrFqdn `json:"appServerInst,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	// Contains the Identifier of the selected Network Slice instance
	NsiId string `json:"nsiId,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio int32 `json:"ratio,omitempty"`

	RatFreq RatFreqInformation `json:"ratFreq,omitempty"`
}
