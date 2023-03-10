/*
 * 3GPP 5GC NRM
 *
 * OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_5GcNrm

type FiveQiCharacteristicsSingleAllOf struct {

	FiveQIValue int32 `json:"fiveQIValue,omitempty"`

	ResourceType string `json:"resourceType,omitempty"`

	PriorityLevel int32 `json:"priorityLevel,omitempty"`

	PacketDelayBudget int32 `json:"packetDelayBudget,omitempty"`

	PacketErrorRate PacketErrorRate `json:"packetErrorRate,omitempty"`

	AveragingWindow int32 `json:"averagingWindow,omitempty"`

	MaximumDataBurstVolume int32 `json:"maximumDataBurstVolume,omitempty"`
}
