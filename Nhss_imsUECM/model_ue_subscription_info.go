/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsUECM

// UeSubscriptionInfo - Subscription information of the UE for the SIP Registration State event
type UeSubscriptionInfo struct {

	CallIdSipHeader string `json:"callIdSipHeader"`

	FromSipHeader string `json:"fromSipHeader"`

	ToSipHeader string `json:"toSipHeader"`

	RecordRoute string `json:"recordRoute"`

	Contact string `json:"contact"`
}
