/*
 * Fault Supervision MnS
 *
 * OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_FaultMnS

type NotifyAckStateChanged struct {

	Href string `json:"href"`

	NotificationId int32 `json:"notificationId"`

	NotificationType NotificationType `json:"notificationType"`

	EventTime time.Time `json:"eventTime"`

	SystemDN string `json:"systemDN"`

	AlarmId string `json:"alarmId"`

	AlarmType AlarmType `json:"alarmType"`

	ProbableCause ProbableCause `json:"probableCause"`

	PerceivedSeverity PerceivedSeverity `json:"perceivedSeverity"`

	AckState AckState `json:"ackState"`

	AckUserId string `json:"ackUserId"`

	AckSystemId string `json:"ackSystemId,omitempty"`
}