/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_MLModelProvision

type UmtTime struct {

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDay string `json:"timeOfDay"`

	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays  shall be indicated with the next higher numbers. 7 shall indicate Sunday. 
	DayOfWeek int32 `json:"dayOfWeek"`
}
