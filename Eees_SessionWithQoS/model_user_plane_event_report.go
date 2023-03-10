/*
 * EES Session with QoS API
 *
 * API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_SessionWithQoS

// UserPlaneEventReport - Represents an event report for user plane.
type UserPlaneEventReport struct {

	Event UserPlaneEvent `json:"event"`

	AccumulatedUsage AccumulatedUsage `json:"accumulatedUsage,omitempty"`

	// Identifies the affected flows that were sent during event subscription. It might be omitted when the reported event applies to all the flows sent during the subscription. 
	FlowIds []int32 `json:"flowIds,omitempty"`

	// The currently applied QoS reference. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION.
	AppliedQosRef string `json:"appliedQosRef,omitempty"`

	// Contains the QoS Monitoring Reporting information
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`
}
