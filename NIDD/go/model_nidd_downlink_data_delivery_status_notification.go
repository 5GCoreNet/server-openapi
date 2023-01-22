/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDD

import (
	"time"
)

// NiddDownlinkDataDeliveryStatusNotification - Represents the delivery status of a specific NIDD downlink data delivery.
type NiddDownlinkDataDeliveryStatusNotification struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NiddDownlinkDataTransfer string `json:"niddDownlinkDataTransfer"`

	DeliveryStatus DeliveryStatus `json:"deliveryStatus"`

	// string with format \"date-time\" as defined in OpenAPI.
	RequestedRetransmissionTime time.Time `json:"requestedRetransmissionTime,omitempty"`
}