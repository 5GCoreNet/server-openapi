/*
 * Nhss_UECM
 *
 * HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_UECM

// DeregistrationRequest - It represents the request body of the deregistration request sent by UDM to HSS and contains the IMSI of the UE, the deregistration reason, etc.
type DeregistrationRequest struct {

	Imsi string `json:"imsi"`

	DeregReason DeregistrationReason `json:"deregReason"`

	Guami Guami `json:"guami,omitempty"`
}
