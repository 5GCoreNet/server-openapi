/*
 * MSGG_N3GDelivery
 *
 * API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MSGG_N3GDelivery

// Address - Contains the Message type data
type Address struct {

	AddrType AddressType `json:"addrType"`

	Addr string `json:"addr"`
}
