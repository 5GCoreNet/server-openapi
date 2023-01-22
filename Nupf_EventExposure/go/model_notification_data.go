/*
 * UPF Event Exposure Service
 *
 * UPF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nupf_EventExposure

// NotificationData - the list of NotificationItems
type NotificationData struct {

	NotificationItems []NotificationItem `json:"notificationItems"`

	CorrelationId string `json:"correlationId,omitempty"`
}
