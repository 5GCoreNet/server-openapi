/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AsSessionWithQoS

type UserPlaneEventAnyOf string

// List of UserPlaneEventAnyOf
const (
	SESSION_TERMINATION UserPlaneEventAnyOf = "SESSION_TERMINATION"
	LOSS_OF_BEARER UserPlaneEventAnyOf = "LOSS_OF_BEARER"
	RECOVERY_OF_BEARER UserPlaneEventAnyOf = "RECOVERY_OF_BEARER"
	RELEASE_OF_BEARER UserPlaneEventAnyOf = "RELEASE_OF_BEARER"
	USAGE_REPORT UserPlaneEventAnyOf = "USAGE_REPORT"
	FAILED_RESOURCES_ALLOCATION UserPlaneEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	QOS_GUARANTEED UserPlaneEventAnyOf = "QOS_GUARANTEED"
	QOS_NOT_GUARANTEED UserPlaneEventAnyOf = "QOS_NOT_GUARANTEED"
	QOS_MONITORING UserPlaneEventAnyOf = "QOS_MONITORING"
	SUCCESSFUL_RESOURCES_ALLOCATION UserPlaneEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
	ACCESS_TYPE_CHANGE UserPlaneEventAnyOf = "ACCESS_TYPE_CHANGE"
	PLMN_CHG UserPlaneEventAnyOf = "PLMN_CHG"
)
