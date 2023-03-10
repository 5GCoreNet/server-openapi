/*
 * Naf_Authentication
 *
 * AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Naf_Authentication

// ReauthRevokeNotify - UAV related notification
type ReauthRevokeNotify struct {

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`

	ServiceLevelId string `json:"serviceLevelId"`

	NotifyCorrId string `json:"notifyCorrId,omitempty"`

	AuthMsg string `json:"authMsg,omitempty"`

	NotifyType NotifyType `json:"notifyType"`

	IpAddr IpAddr `json:"ipAddr,omitempty"`
}
