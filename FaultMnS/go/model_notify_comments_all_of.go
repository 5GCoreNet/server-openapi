/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package FaultMnS

type NotifyCommentsAllOf struct {

	AlarmId string `json:"alarmId"`

	AlarmType AlarmType `json:"alarmType"`

	ProbableCause ProbableCause `json:"probableCause"`

	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`

	// Collection of comments. The comment identifiers are allocated by the MnS producer and used as key in the map.
	Comments map[string]Comment `json:"comments"`
}
