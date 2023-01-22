/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nchf_ConvergedCharging

type NspaContainerInformation struct {

	Latency int32 `json:"latency,omitempty"`

	Throughput Throughput `json:"throughput,omitempty"`

	MaximumPacketLossRate string `json:"maximumPacketLossRate,omitempty"`

	ServiceExperienceStatisticsData ServiceExperienceInfo `json:"serviceExperienceStatisticsData,omitempty"`

	TheNumberOfPDUSessions int32 `json:"theNumberOfPDUSessions,omitempty"`

	TheNumberOfRegisteredSubscribers int32 `json:"theNumberOfRegisteredSubscribers,omitempty"`

	LoadLevel NsiLoadLevelInfo `json:"loadLevel,omitempty"`
}
