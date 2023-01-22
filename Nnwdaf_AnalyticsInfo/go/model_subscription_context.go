/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// SubscriptionContext - Context data related to a created subscription, to be included in notifications sent by NRF 
type SubscriptionContext struct {

	SubscriptionId string `json:"subscriptionId"`

	SubscrCond SubscrCond `json:"subscrCond,omitempty"`
}
