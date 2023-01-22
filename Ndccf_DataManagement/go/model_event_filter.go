/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// EventFilter - Represents event filter information for an event.
type EventFilter struct {

	Gpsis []string `json:"gpsis,omitempty"`

	Supis []string `json:"supis,omitempty"`

	ExterGroupIds []string `json:"exterGroupIds,omitempty"`

	InterGroupIds []string `json:"interGroupIds,omitempty"`

	AnyUeInd bool `json:"anyUeInd,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	LocArea LocationArea5G `json:"locArea,omitempty"`

	CollAttrs []CollectiveBehaviourFilter `json:"collAttrs,omitempty"`
}
