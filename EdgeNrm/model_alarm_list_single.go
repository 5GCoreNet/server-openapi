/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EdgeNrm

type AlarmListSingle struct {

	Id *string `json:"id"`

	ObjectClass string `json:"objectClass,omitempty"`

	ObjectInstance string `json:"objectInstance,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`

	Attributes AlarmListSingleAllOfAttributes `json:"attributes,omitempty"`
}
