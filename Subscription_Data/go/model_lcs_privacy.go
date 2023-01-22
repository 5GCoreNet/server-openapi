/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type LcsPrivacy struct {

	AfInstanceId string `json:"afInstanceId,omitempty"`

	ReferenceId int32 `json:"referenceId,omitempty"`

	Lpi Lpi `json:"lpi,omitempty"`

	// String uniquely identifying MTC provider information.
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
}
