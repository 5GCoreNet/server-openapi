/*
 * 3gpp-bdt
 *
 * API for BDT resouce management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ResourceManagementOfBdt

// BdtPatch - Represents a Background Data Transfer subscription modification request.
type BdtPatch struct {

	// Identity of the selected background data transfer policy.
	SelectedPolicy int32 `json:"selectedPolicy"`

	// Indicates whether the BDT warning notification is enabled (true) or not (false). 
	WarnNotifEnabled bool `json:"warnNotifEnabled,omitempty"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}
