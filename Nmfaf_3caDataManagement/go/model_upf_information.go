/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// UpfInformation - Represents the ID/address/FQDN of the UPF.
type UpfInformation struct {

	UpfId string `json:"upfId,omitempty"`

	UpfAddr AddrFqdn `json:"upfAddr,omitempty"`
}
