/*
 * MDA Report
 *
 * OAS 3.0.1 specification of the MDA Report © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MdaReport

type VsDataContainerSingle struct {

	Id string `json:"id,omitempty"`

	Attributes VsDataContainerSingleAttributes `json:"attributes,omitempty"`

	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
}
