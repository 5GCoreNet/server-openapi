/*
 * MSGS_MSGDelivery
 *
 * API for MSGG MSGin5G Server Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MSGS_MSGDelivery

// MessageSegmentParameters - Contains the message segment parameters data
type MessageSegmentParameters struct {

	SegId string `json:"segId,omitempty"`

	TotalSegCount int32 `json:"totalSegCount,omitempty"`

	SegNumb int32 `json:"segNumb,omitempty"`

	LastSegFlag bool `json:"lastSegFlag,omitempty"`
}
