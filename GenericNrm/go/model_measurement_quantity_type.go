/*
 * Generic NRM
 *
 * OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package GenericNrm
// MeasurementQuantityType : See details in 3GPP TS 32.422 clause 5.10.15.
type MeasurementQuantityType string

// List of MeasurementQuantityType
const (
	CPICH_EC_NO MeasurementQuantityType = "CPICH_EcNo"
	CPICH_RSCP MeasurementQuantityType = "CPICH_RSCP"
	PATH_LOSS MeasurementQuantityType = "PathLoss"
)