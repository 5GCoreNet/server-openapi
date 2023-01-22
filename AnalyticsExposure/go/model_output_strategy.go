/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

// OutputStrategy - Possible values are: - BINARY: Indicates that the analytics shall only be reported when the requested level of accuracy is reached within a cycle of periodic notification. - GRADIENT: Indicates that the analytics shall be reported according with the periodicity irrespective of whether the requested level of accuracy has been reached or not. 
type OutputStrategy struct {
}
