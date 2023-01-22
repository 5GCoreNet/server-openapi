/*
 * SS_NetworkResourceMonitoring
 *
 * API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_NetworkResourceMonitoring

type MeasurementDataTypeAnyOf string

// List of MeasurementDataTypeAnyOf
const (
	DL_DELAY MeasurementDataTypeAnyOf = "DL_DELAY"
	UL_DELAY MeasurementDataTypeAnyOf = "UL_DELAY"
	RT_DELAY MeasurementDataTypeAnyOf = "RT_DELAY"
	AVG_PLR MeasurementDataTypeAnyOf = "AVG_PLR"
	AVG_DATA_RATE MeasurementDataTypeAnyOf = "AVG_DATA_RATE"
	MAX_DATA_RATE MeasurementDataTypeAnyOf = "MAX_DATA_RATE"
	AVG_DL_TRAFFIC_VOLUME MeasurementDataTypeAnyOf = "AVG_DL_TRAFFIC_VOLUME"
	AVG_UL_TRAFFIC_VOLUME MeasurementDataTypeAnyOf = "AVG_UL_TRAFFIC_VOLUME"
)
