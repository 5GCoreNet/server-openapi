/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssf_NSSelection

// SubscribedSnssai - Contains the subscribed S-NSSAI
type SubscribedSnssai struct {

	SubscribedSnssai Snssai `json:"subscribedSnssai"`

	DefaultIndication bool `json:"defaultIndication,omitempty"`

	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
}