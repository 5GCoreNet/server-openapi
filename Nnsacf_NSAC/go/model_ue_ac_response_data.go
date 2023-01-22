/*
 * Nnsacf_NSAC
 *
 * Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnsacf_NSAC

type UeAcResponseData struct {

	// A map (list of key-value pairs) where the key of the map shall be UE's supi
	AcuFailureList map[string][]AcuFailureItem `json:"acuFailureList,omitempty"`
}
