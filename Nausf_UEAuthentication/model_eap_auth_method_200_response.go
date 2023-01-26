/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nausf_UEAuthentication

type EapAuthMethod200Response struct {

	// contains an EAP packet
	EapPayload *string `json:"eapPayload"`

	// URI : /{eapSessionUri}, a map(list of key-value pairs) where member serves as key
	Links map[string]LinksValueSchema `json:"_links"`
}