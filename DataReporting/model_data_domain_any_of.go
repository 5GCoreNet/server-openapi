/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReporting

type DataDomainAnyOf string

// List of DataDomainAnyOf
const (
	SERVICE_EXPERIENCE DataDomainAnyOf = "SERVICE_EXPERIENCE"
	LOCATION DataDomainAnyOf = "LOCATION"
	COMMUNICATION DataDomainAnyOf = "COMMUNICATION"
	PERFORMANCE DataDomainAnyOf = "PERFORMANCE"
	APPLICATION_SPECIFIC DataDomainAnyOf = "APPLICATION_SPECIFIC"
	MS_ACCESS_ACTIVITY DataDomainAnyOf = "MS_ACCESS_ACTIVITY"
	PLANNED_TRIPS DataDomainAnyOf = "PLANNED_TRIPS"
)
