/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

type MediaStreamingAccessRecordAllOf struct {

	MediaStreamHandlerEndpointAddress EndpointAddress `json:"mediaStreamHandlerEndpointAddress"`

	ApplicationServerEndpointAddress EndpointAddress `json:"applicationServerEndpointAddress"`

	SessionIdentifier string `json:"sessionIdentifier,omitempty"`

	RequestMessage MediaStreamingAccessRecordAllOfRequestMessage `json:"requestMessage"`

	CacheStatus CacheStatus `json:"cacheStatus,omitempty"`

	ResponseMessage MediaStreamingAccessRecordAllOfResponseMessage `json:"responseMessage"`

	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32 `json:"processingLatency"`

	ConnectionMetrics MediaStreamingAccessRecordAllOfConnectionMetrics `json:"connectionMetrics,omitempty"`
}
