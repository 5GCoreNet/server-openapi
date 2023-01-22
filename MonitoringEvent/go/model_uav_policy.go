/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MonitoringEvent

// UavPolicy - Represents the policy information included in the UAV presence monitoring request.
type UavPolicy struct {

	UavMoveInd bool `json:"uavMoveInd"`

	RevokeInd bool `json:"revokeInd"`
}
