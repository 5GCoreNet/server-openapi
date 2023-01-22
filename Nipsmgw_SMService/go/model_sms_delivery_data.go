/*
 * Nipsmgw_SMService Service API
 *
 * IP-SM-GW SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nipsmgw_SMService

// SmsDeliveryData - Information within response message invoking MtForwardSm service operation, for delivering MT SMS Delivery Report. 
type SmsDeliveryData struct {

	SmsPayload RefToBinaryData `json:"smsPayload"`
}
