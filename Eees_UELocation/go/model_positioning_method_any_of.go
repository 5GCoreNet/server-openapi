/*
 * EES UE Location Information_API
 *
 * API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_UELocation

type PositioningMethodAnyOf string

// List of PositioningMethodAnyOf
const (
	CELLID PositioningMethodAnyOf = "CELLID"
	ECID PositioningMethodAnyOf = "ECID"
	OTDOA PositioningMethodAnyOf = "OTDOA"
	BAROMETRIC_PRESSURE PositioningMethodAnyOf = "BAROMETRIC_PRESSURE"
	WLAN PositioningMethodAnyOf = "WLAN"
	BLUETOOTH PositioningMethodAnyOf = "BLUETOOTH"
	MBS PositioningMethodAnyOf = "MBS"
	MOTION_SENSOR PositioningMethodAnyOf = "MOTION_SENSOR"
	DL_TDOA PositioningMethodAnyOf = "DL_TDOA"
	DL_AOD PositioningMethodAnyOf = "DL_AOD"
	MULTI_RTT PositioningMethodAnyOf = "MULTI-RTT"
	NR_ECID PositioningMethodAnyOf = "NR_ECID"
	UL_TDOA PositioningMethodAnyOf = "UL_TDOA"
	UL_AOA PositioningMethodAnyOf = "UL_AOA"
	NETWORK_SPECIFIC PositioningMethodAnyOf = "NETWORK_SPECIFIC"
)
