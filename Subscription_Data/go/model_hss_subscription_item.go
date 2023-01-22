/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

// HssSubscriptionItem - Contains info about a single HSS event subscription
type HssSubscriptionItem struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	HssInstanceId string `json:"hssInstanceId"`

	// String providing an URI formatted according to RFC 3986.
	SubscriptionId string `json:"subscriptionId"`

	ContextInfo ContextInfo `json:"contextInfo,omitempty"`
}
