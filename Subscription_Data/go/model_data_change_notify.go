/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

// DataChangeNotify - Container for data which have changed and notification was requested when changed.
type DataChangeNotify struct {

	OriginalCallbackReference []string `json:"originalCallbackReference,omitempty"`

	// String represents the SUPI or GPSI
	UeId string `json:"ueId,omitempty"`

	NotifyItems []NotifyItem `json:"notifyItems,omitempty"`

	SdmSubscription SdmSubscription `json:"sdmSubscription,omitempty"`

	AdditionalSdmSubscriptions []SdmSubscription `json:"additionalSdmSubscriptions,omitempty"`

	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
}
