/*
 * Namf_MT
 *
 * AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Namf_MT

// UeInfo - list of UEs requested to be made reachable for the MBS Session
type UeInfo struct {

	UeList []string `json:"ueList"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId,omitempty"`
}
