/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EdgeNrm

import (
	"time"
)

// FileDownloadJobProcessMonitor - This data type is the \"ProcessMonitor\" data type with specialisations for usage in the \"FileDownloadJob\".
type FileDownloadJobProcessMonitor struct {

	JobId string `json:"jobId,omitempty"`

	Status string `json:"status,omitempty"`

	ProgressPercentage int32 `json:"progressPercentage,omitempty"`

	ProgressStateInfo string `json:"progressStateInfo,omitempty"`

	ResultStateInfo FileDownloadJobProcessMonitorResultStateInfo `json:"resultStateInfo,omitempty"`

	StartTime time.Time `json:"startTime,omitempty"`

	EndTime time.Time `json:"endTime,omitempty"`

	Timer int32 `json:"timer,omitempty"`
}
