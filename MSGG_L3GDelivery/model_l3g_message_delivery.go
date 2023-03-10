/*
 * MSGG_L3GDelivery
 *
 * API for MSGG L3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_L3GDelivery

// L3gMessageDelivery - Contains the L3G message delivery data
type L3gMessageDelivery struct {

	OriAddr Address `json:"oriAddr"`

	DestAddr Address `json:"destAddr"`

	AppId string `json:"appId,omitempty"`

	MsgId string `json:"msgId"`

	DelivStReqInd bool `json:"delivStReqInd,omitempty"`

	Payload string `json:"payload,omitempty"`

	SegInd bool `json:"segInd,omitempty"`

	SegParams MessageSegmentParameters `json:"segParams,omitempty"`
}
