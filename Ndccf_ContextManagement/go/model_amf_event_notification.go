/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_ContextManagement

// AmfEventNotification - Data within a AMF Event Notification request
type AmfEventNotification struct {

	NotifyCorrelationId string `json:"notifyCorrelationId,omitempty"`

	SubsChangeNotifyCorrelationId string `json:"subsChangeNotifyCorrelationId,omitempty"`

	ReportList []AmfEventReport `json:"reportList,omitempty"`

	EventSubsSyncInfo AmfEventSubsSyncInfo `json:"eventSubsSyncInfo,omitempty"`
}
