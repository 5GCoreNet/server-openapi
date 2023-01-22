/*
 * Nnef_PFDmanagement Service API
 *
 * Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_PFDmanagement

// NotificationPush - Represents the information to be used by the NF service consumer to retrieve the PFDs and/or remove the PFDs of the applicable application identifier(s). 
type NotificationPush struct {

	AppIds []string `json:"appIds"`

	// indicating a time in seconds.
	AllowedDelay int32 `json:"allowedDelay,omitempty"`

	PfdOp PfdOperation `json:"pfdOp,omitempty"`
}