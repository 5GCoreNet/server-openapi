/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AnalyticsExposure

// NsiIdInfo - Represents the S-NSSAI and the optionally associated Network Slice Instance(s).
type NsiIdInfo struct {

	Snssai Snssai `json:"snssai"`

	NsiIds []string `json:"nsiIds,omitempty"`
}
