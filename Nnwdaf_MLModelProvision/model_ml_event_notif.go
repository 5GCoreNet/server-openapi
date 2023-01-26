/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_MLModelProvision

// MlEventNotif - Represents a notification related to a single event that occurred.
type MlEventNotif struct {

	Event NwdafEvent `json:"event"`

	NotifCorreId string `json:"notifCorreId,omitempty"`

	MLFileAddr MlModelAddr `json:"mLFileAddr"`

	ValidityPeriod TimeWindow `json:"validityPeriod,omitempty"`

	SpatialValidity NetworkAreaInfo `json:"spatialValidity,omitempty"`
}