/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_Communication

// ConfiguredSnssai - Contains the configured S-NSSAI(s) authorized by the NSSF in the serving PLMN and optional mapped home S-NSSAI
type ConfiguredSnssai struct {

	ConfiguredSnssai Snssai `json:"configuredSnssai"`

	MappedHomeSnssai Snssai `json:"mappedHomeSnssai,omitempty"`
}
