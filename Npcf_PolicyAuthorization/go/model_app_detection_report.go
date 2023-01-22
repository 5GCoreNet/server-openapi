/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Npcf_PolicyAuthorization

// AppDetectionReport - Indicates the start or stop of the detected application traffic and the application identifier of the detected application traffic.
type AppDetectionReport struct {

	AdNotifType AppDetectionNotifType `json:"adNotifType"`

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId"`
}
