/*
 * 3gpp-network-status-reporting
 *
 * API for reporting network status.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ReportingNetworkStatus

import (
	"time"
)

// NetStatusRepSubsPatch - Represents the parameters to request the modification of network status reporting subscription.
type NetStatusRepSubsPatch struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty"`

	LocationArea LocationArea `json:"locationArea,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI with \"nullable=true\" property.
	TimeDuration *time.Time `json:"timeDuration,omitempty"`

	ThresholdValues []int32 `json:"thresholdValues,omitempty"`

	ThresholdTypes []CongestionType `json:"thresholdTypes,omitempty"`
}
