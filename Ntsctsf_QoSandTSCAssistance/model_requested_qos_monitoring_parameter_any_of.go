/*
 * Ntsctsf_QoSandTSCAssistance Service API
 *
 * TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ntsctsf_QoSandTSCAssistance

type RequestedQosMonitoringParameterAnyOf string

// List of RequestedQosMonitoringParameterAnyOf
const (
	DOWNLINK RequestedQosMonitoringParameterAnyOf = "DOWNLINK"
	UPLINK RequestedQosMonitoringParameterAnyOf = "UPLINK"
	ROUND_TRIP RequestedQosMonitoringParameterAnyOf = "ROUND_TRIP"
)
