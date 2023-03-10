/*
 * UAE Server Real-time UAV Status Service
 *
 * UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_RealtimeUAVStatus

import (
	"time"
)

// UavNetConnStatus - Represents UAV network connection status information.
type UavNetConnStatus struct {

	StatusInfo MonitoringType `json:"statusInfo"`

	// string with format \"date-time\" as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
}
