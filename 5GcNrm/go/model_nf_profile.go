/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5GcNrm

// NfProfile - NF profile stored in NRF, defined in TS 29.510
type NfProfile struct {

	// uuid of NF instance
	NFInstanceId string `json:"nFInstanceId,omitempty"`

	NFType NfType `json:"nFType,omitempty"`

	NFStatus NfStatus `json:"nFStatus,omitempty"`

	Plmn PlmnId `json:"plmn,omitempty"`

	SNssais Snssai `json:"sNssais,omitempty"`

	Fqdn string `json:"fqdn,omitempty"`

	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty"`

	NfServices []NfService `json:"nfServices,omitempty"`
}
