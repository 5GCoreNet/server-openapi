/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm
// SupportedDataSet : any of enumrated value
type SupportedDataSet string

// List of SupportedDataSet
const (
	SUBSCRIPTION SupportedDataSet = "SUBSCRIPTION"
	POLICY SupportedDataSet = "POLICY"
	EXPOSURE SupportedDataSet = "EXPOSURE"
	APPLICATION SupportedDataSet = "APPLICATION"
	A_PFD SupportedDataSet = "A_PFD"
	A_AFTI SupportedDataSet = "A_AFTI"
	A_IPTV SupportedDataSet = "A_IPTV"
	A_BDT SupportedDataSet = "A_BDT"
	A_SPD SupportedDataSet = "A_SPD"
	A_EASD SupportedDataSet = "A_EASD"
	A_AMI SupportedDataSet = "A_AMI"
	P_UE SupportedDataSet = "P_UE"
	P_SCD SupportedDataSet = "P_SCD"
	P_BDT SupportedDataSet = "P_BDT"
	P_PLMNUE SupportedDataSet = "P_PLMNUE"
	P_NSSCD SupportedDataSet = "P_NSSCD"
)