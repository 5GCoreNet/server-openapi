/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AiMlNrm

type ManagementDataCollectionSingleAllOfAttributes struct {

	ManagementData ManagementData `json:"managementData,omitempty"`

	TargetNodeFilter NodeFilter `json:"targetNodeFilter,omitempty"`

	CollectionTimeWindow TimeWindow `json:"collectionTimeWindow,omitempty"`

	ReportingCtrl string `json:"reportingCtrl,omitempty"`

	DataScope string `json:"dataScope,omitempty"`
}
