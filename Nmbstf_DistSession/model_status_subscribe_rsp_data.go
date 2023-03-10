/*
 * Nmbstf-distsession
 *
 * MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbstf_DistSession

// StatusSubscribeRspData - Data within StatusSubscribe Response
type StatusSubscribeRspData struct {

	Subscription DistSessionSubscription `json:"subscription"`

	ReportList DistSessionEventReportList `json:"reportList,omitempty"`
}
