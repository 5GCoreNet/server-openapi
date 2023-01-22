/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5GcNrm

// ListOfInterfacesType - The interfaces to be recorded in the Network Element. See 3GPP TS 32.422 clause 5.5 for additional details.
type ListOfInterfacesType struct {

	MSCServerInterfaces []string `json:"MSCServerInterfaces,omitempty"`

	MGWInterfaces []string `json:"MGWInterfaces,omitempty"`

	RNCInterfaces []string `json:"RNCInterfaces,omitempty"`

	SGSNInterfaces []string `json:"SGSNInterfaces,omitempty"`

	GGSNInterfaces []string `json:"GGSNInterfaces,omitempty"`

	SCSCFInterfaces []string `json:"S-CSCFInterfaces,omitempty"`

	PCSCFInterfaces []string `json:"P-CSCFInterfaces,omitempty"`

	ICSCFInterfaces []string `json:"I-CSCFInterfaces,omitempty"`

	MRFCInterfaces []string `json:"MRFCInterfaces,omitempty"`

	MGCFInterfaces []string `json:"MGCFInterfaces,omitempty"`

	IBCFInterfaces []string `json:"IBCFInterfaces,omitempty"`

	ECSCFInterfaces []string `json:"E-CSCFInterfaces,omitempty"`

	BGCFInterfaces []string `json:"BGCFInterfaces,omitempty"`

	ASInterfaces []string `json:"ASInterfaces,omitempty"`

	HSSInterfaces []string `json:"HSSInterfaces,omitempty"`

	EIRInterfaces []string `json:"EIRInterfaces,omitempty"`

	BMSCInterfaces []string `json:"BM-SCInterfaces,omitempty"`

	MMEInterfaces []string `json:"MMEInterfaces,omitempty"`

	SGWInterfaces []string `json:"SGWInterfaces,omitempty"`

	PDNGWInterfaces []string `json:"PDN_GWInterfaces,omitempty"`

	ENBInterfaces []string `json:"eNBInterfaces,omitempty"`

	EnGNBInterfaces []string `json:"en-gNBInterfaces,omitempty"`

	AMFInterfaces []string `json:"AMFInterfaces,omitempty"`

	AUSFInterfaces []string `json:"AUSFInterfaces,omitempty"`

	NEFInterfaces []string `json:"NEFInterfaces,omitempty"`

	NRFInterfaces []string `json:"NRFInterfaces,omitempty"`

	NSSFInterfaces []string `json:"NSSFInterfaces,omitempty"`

	PCFInterfaces []string `json:"PCFInterfaces,omitempty"`

	SMFInterfaces []string `json:"SMFInterfaces,omitempty"`

	SMSFInterfaces []string `json:"SMSFInterfaces,omitempty"`

	UDMInterfaces []string `json:"UDMInterfaces,omitempty"`

	UPFInterfaces []string `json:"UPFInterfaces,omitempty"`

	NgENBInterfaces []string `json:"ng-eNBInterfaces,omitempty"`

	GNBCUCPInterfaces []string `json:"gNB-CU-CPInterfaces,omitempty"`

	GNBCUUPInterfaces []string `json:"gNB-CU-UPInterfaces,omitempty"`

	GNBDUInterfaces []string `json:"gNB-DUInterfaces,omitempty"`
}
