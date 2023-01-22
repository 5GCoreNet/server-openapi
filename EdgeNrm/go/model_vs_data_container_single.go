/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EdgeNrm

type VsDataContainerSingle struct {

	Id string `json:"id,omitempty"`

	Attributes VsDataContainerSingleAttributes `json:"attributes,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
}