/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

// NotificationData - Data sent in notifications from NRF to subscribed NF Instances
type NotificationData struct {

	Event NotificationEventType `json:"event"`

	// String providing an URI formatted according to RFC 3986.
	NfInstanceUri string `json:"nfInstanceUri"`

	NfProfile NfProfile `json:"nfProfile,omitempty"`

	ProfileChanges []ChangeItem `json:"profileChanges,omitempty"`

	ConditionEvent ConditionEventType `json:"conditionEvent,omitempty"`

	SubscriptionContext SubscriptionContext `json:"subscriptionContext,omitempty"`
}
