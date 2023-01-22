/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_DataManagement

// DataSubscription - Contains a data specification.
type DataSubscription struct {

	AmfDataSub AmfEventSubscription `json:"amfDataSub,omitempty"`

	SmfDataSub NsmfEventExposure `json:"smfDataSub,omitempty"`

	UdmDataSub EeSubscription `json:"udmDataSub,omitempty"`

	AfDataSub AfEventExposureSubsc `json:"afDataSub,omitempty"`

	NefDataSub NefEventExposureSubsc `json:"nefDataSub,omitempty"`

	NrfDataSub SubscriptionData `json:"nrfDataSub,omitempty"`

	NsacfDataSub SacEventSubscription `json:"nsacfDataSub,omitempty"`
}
