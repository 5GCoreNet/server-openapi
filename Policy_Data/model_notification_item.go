/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

// NotificationItem - Identifies a data change notification when the change occurs in a fragment (subset of resource data) of a given resource. 
type NotificationItem struct {

	// String providing an URI formatted according to RFC 3986.
	ResourceId string `json:"resourceId"`

	NotifItems []UpdatedItem `json:"notifItems"`
}