/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N32_Handshake

// IpxProviderSecInfo - Defines the security information list of an IPX
type IpxProviderSecInfo struct {

	// Fully Qualified Domain Name
	IpxProviderId string `json:"ipxProviderId"`

	RawPublicKeyList []string `json:"rawPublicKeyList,omitempty"`

	CertificateList []string `json:"certificateList,omitempty"`
}
