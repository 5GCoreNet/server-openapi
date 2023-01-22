/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// NnwdafEventsSubscriptionNotification - Represents an Individual NWDAF Event Subscription Notification resource.
type NnwdafEventsSubscriptionNotification struct {

	// Notifications about Individual Events
	EventNotifications []EventNotification `json:"eventNotifications,omitempty"`

	// String identifying a subscription to the Nnwdaf_EventsSubscription Service
	SubscriptionId string `json:"subscriptionId"`

	// Notification correlation identifier.
	NotifCorrId string `json:"notifCorrId,omitempty"`

	// Subscription ID which was allocated by the source NWDAF. This parameter shall be present if the notification is for informing the assignment of a new Subscription Id by the target NWDAF. 
	OldSubscriptionId string `json:"oldSubscriptionId,omitempty"`
}
