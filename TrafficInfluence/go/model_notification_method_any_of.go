/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TrafficInfluence

type NotificationMethodAnyOf string

// List of NotificationMethodAnyOf
const (
	PERIODIC NotificationMethodAnyOf = "PERIODIC"
	ONE_TIME NotificationMethodAnyOf = "ONE_TIME"
	ON_EVENT_DETECTION NotificationMethodAnyOf = "ON_EVENT_DETECTION"
)
