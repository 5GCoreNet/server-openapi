/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsUECM

// RestorationInfo - The information relevant to a specific registration required for an S-CSCF to handle the requests for a user 
type RestorationInfo struct {

	Path string `json:"path"`

	Contact string `json:"contact"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	InitialCSeqSequenceNumber int32 `json:"initialCSeqSequenceNumber,omitempty"`

	CallIdSipHeader string `json:"callIdSipHeader,omitempty"`

	UesubscriptionInfo UeSubscriptionInfo `json:"uesubscriptionInfo,omitempty"`

	PcscfSubscriptionInfo PcscfSubscriptionInfo `json:"pcscfSubscriptionInfo,omitempty"`

	// A map (list of key-value pairs where subscriptionId serves as key) of ImsSdmSubscription 
	ImsSdmSubscriptions map[string]ImsSdmSubscription `json:"imsSdmSubscriptions,omitempty"`
}
