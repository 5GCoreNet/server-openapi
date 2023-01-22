/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudsf_DataRepository

// SubscriptionFilter - A subscription filter
type SubscriptionFilter struct {

	// list of resources applicable to the subscription
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`

	// list of resources applicable to the subscription
	Operations []RecordOperation `json:"operations,omitempty"`
}
