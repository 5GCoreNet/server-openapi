/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SliceNrm

type NtfSubscriptionControlSingle struct {

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes NtfSubscriptionControlSingleAllOfAttributes `json:"attributes,omitempty"`

	HeartbeatControl HeartbeatControlSingle `json:"HeartbeatControl,omitempty"`
}
