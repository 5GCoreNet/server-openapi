/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

import (
	"time"
)

type UsedUnitContainer struct {

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	ServiceId int32 `json:"serviceId,omitempty"`

	QuotaManagementIndicator QuotaManagementIndicator `json:"quotaManagementIndicator,omitempty"`

	Triggers []Trigger `json:"triggers,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TriggerTimestamp time.Time `json:"triggerTimestamp,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Time int32 `json:"time,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	TotalVolume int32 `json:"totalVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UplinkVolume int32 `json:"uplinkVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	DownlinkVolume int32 `json:"downlinkVolume,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	ServiceSpecificUnits int32 `json:"serviceSpecificUnits,omitempty"`

	EventTimeStamps []time.Time `json:"eventTimeStamps,omitempty"`

	LocalSequenceNumber int32 `json:"localSequenceNumber"`

	PDUContainerInformation PduContainerInformation `json:"pDUContainerInformation,omitempty"`

	NSPAContainerInformation NspaContainerInformation `json:"nSPAContainerInformation,omitempty"`

	PC5ContainerInformation Pc5ContainerInformation `json:"pC5ContainerInformation,omitempty"`
}
