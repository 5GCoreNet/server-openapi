/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

// ApiInvokerPolicy - Represents the policy of an API Invoker.
type ApiInvokerPolicy struct {

	// API invoker ID assigned by the CAPIF core function
	ApiInvokerId string `json:"apiInvokerId"`

	// Total number of invocations allowed on the service API by the API invoker.
	AllowedTotalInvocations int32 `json:"allowedTotalInvocations,omitempty"`

	// Invocations per second allowed on the service API by the API invoker.
	AllowedInvocationsPerSecond int32 `json:"allowedInvocationsPerSecond,omitempty"`

	// The time ranges during which the invocations are allowed on the service API by the API invoker. 
	AllowedInvocationTimeRangeList []TimeRangeList `json:"allowedInvocationTimeRangeList,omitempty"`
}
