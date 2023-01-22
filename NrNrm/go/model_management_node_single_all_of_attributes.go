/*
 * NR NRM
 *
 * OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NrNrm

type ManagementNodeSingleAllOfAttributes struct {

	UserLabel string `json:"userLabel,omitempty"`

	ManagedElements []string `json:"managedElements,omitempty"`

	VendorName string `json:"vendorName,omitempty"`

	UserDefinedState string `json:"userDefinedState,omitempty"`

	LocationName string `json:"locationName,omitempty"`

	SwVersion string `json:"swVersion,omitempty"`
}
