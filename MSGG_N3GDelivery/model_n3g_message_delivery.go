/*
 * MSGG_N3GDelivery
 *
 * API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_N3GDelivery

// N3gMessageDelivery - N3G message delivery data
type N3gMessageDelivery struct {

	OriAddr Address `json:"oriAddr"`

	DestAddr Address `json:"destAddr"`

	AppId string `json:"appId,omitempty"`

	MsgId string `json:"msgId"`

	DelivStReqInd bool `json:"delivStReqInd,omitempty"`

	Payload string `json:"payload,omitempty"`

	SegInd bool `json:"segInd,omitempty"`

	SegParams MessageSegmentParameters `json:"segParams,omitempty"`
}
