/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_SMPolicyControl

// QosNotificationControlInfo - Contains the QoS Notification Control Information.
type QosNotificationControlInfo struct {

	// An array of PCC rule id references to the PCC rules associated with the QoS notification control info.
	RefPccRuleIds []string `json:"refPccRuleIds"`

	NotifType QosNotifType `json:"notifType"`

	// Represents the content version of some content.
	ContVer int32 `json:"contVer,omitempty"`

	AltQosParamId string `json:"altQosParamId,omitempty"`
}
