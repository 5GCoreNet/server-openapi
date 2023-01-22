/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_UEPolicyControl

// Guami - Globally Unique AMF Identifier constructed out of PLMN, Network and AMF identity.
type Guami struct {

	PlmnId PlmnIdNid `json:"plmnId"`

	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).  
	AmfId string `json:"amfId"`
}
