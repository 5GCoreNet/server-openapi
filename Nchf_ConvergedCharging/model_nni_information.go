/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type NniInformation struct {

	SessionDirection NniSessionDirection `json:"sessionDirection,omitempty"`

	NNIType NniType `json:"nNIType,omitempty"`

	RelationshipMode NniRelationshipMode `json:"relationshipMode,omitempty"`

	NeighbourNodeAddress ImsAddress `json:"neighbourNodeAddress,omitempty"`
}
