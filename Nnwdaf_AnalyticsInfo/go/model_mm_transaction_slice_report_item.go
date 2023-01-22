/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

import (
	"time"
)

// MmTransactionSliceReportItem - UE MM Transaction Report Item per Slice
type MmTransactionSliceReportItem struct {

	Snssai Snssai `json:"snssai,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`

	Transactions int32 `json:"transactions"`
}
