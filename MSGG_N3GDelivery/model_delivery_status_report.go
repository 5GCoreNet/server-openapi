/*
 * MSGG_N3GDelivery
 *
 * API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_N3GDelivery

// DeliveryStatusReport - Contains the delivery status report data
type DeliveryStatusReport struct {

	OriAddr Address `json:"oriAddr"`

	DestAddr Address `json:"destAddr"`

	MsgId string `json:"msgId"`

	SecCred string `json:"secCred,omitempty"`

	FailureCause string `json:"failureCause,omitempty"`

	DelivSt ReportDeliveryStatus `json:"delivSt"`
}
