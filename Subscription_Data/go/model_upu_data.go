/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

import (
	"time"
)

// UpuData - Used to store the status of the latest UPU data update.
type UpuData struct {

	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`

	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuXmacIue string `json:"upuXmacIue,omitempty"`

	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue string `json:"upuMacIue,omitempty"`
}
