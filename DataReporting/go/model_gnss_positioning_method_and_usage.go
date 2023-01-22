/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataReporting

// GnssPositioningMethodAndUsage - Indicates the usage of a Global Navigation Satellite System (GNSS) positioning method.
type GnssPositioningMethodAndUsage struct {

	Mode PositioningMode `json:"mode"`

	Gnss GnssId `json:"gnss"`

	Usage Usage `json:"usage"`
}
