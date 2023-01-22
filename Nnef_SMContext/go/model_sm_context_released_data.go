/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_SMContext

// SmContextReleasedData - Successful release of an Individual SM context with information sent to the NF service consumer.
type SmContextReleasedData struct {

	SmallDataRateStatus SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`

	ApnRateStatus ApnRateStatus `json:"apnRateStatus,omitempty"`
}
