/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReporting

type ServiceExperienceRecord struct {

	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`

	ServiceExperienceInfos PerFlowServiceExperienceInfo `json:"serviceExperienceInfos"`
}
