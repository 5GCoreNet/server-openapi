/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// ProSeCapability - Indicate the supported ProSe Capability by the PCF.
type ProSeCapability struct {

	ProseDirectDiscovey bool `json:"proseDirectDiscovey,omitempty"`

	ProseDirectCommunication bool `json:"proseDirectCommunication,omitempty"`

	ProseL2UetoNetworkRelay bool `json:"proseL2UetoNetworkRelay,omitempty"`

	ProseL3UetoNetworkRelay bool `json:"proseL3UetoNetworkRelay,omitempty"`

	ProseL2RemoteUe bool `json:"proseL2RemoteUe,omitempty"`

	ProseL3RemoteUe bool `json:"proseL3RemoteUe,omitempty"`
}
