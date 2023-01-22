/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

type EarlyMediaDescription struct {

	SDPTimeStamps SdpTimeStamps `json:"sDPTimeStamps,omitempty"`

	SDPMediaComponent []SdpMediaComponent `json:"sDPMediaComponent,omitempty"`

	SDPSessionDescription []string `json:"sDPSessionDescription,omitempty"`
}
