/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TimeSyncExposure

type AsTimeResourceAnyOf string

// List of AsTimeResourceAnyOf
const (
	ATOMIC_CLOCK AsTimeResourceAnyOf = "ATOMIC_CLOCK"
	GNSS AsTimeResourceAnyOf = "GNSS"
	TERRESTRIAL_RADIO AsTimeResourceAnyOf = "TERRESTRIAL_RADIO"
	SERIAL_TIME_CODE AsTimeResourceAnyOf = "SERIAL_TIME_CODE"
	PTP AsTimeResourceAnyOf = "PTP"
	NTP AsTimeResourceAnyOf = "NTP"
	HAND_SET AsTimeResourceAnyOf = "HAND_SET"
	INTERNAL_OSCILLATOR AsTimeResourceAnyOf = "INTERNAL_OSCILLATOR"
	OTHER AsTimeResourceAnyOf = "OTHER"
)
