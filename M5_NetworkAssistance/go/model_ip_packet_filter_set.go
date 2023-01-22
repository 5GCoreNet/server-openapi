/*
 * M5_NetworkAssistance
 *
 * 5GMS AF M5 Network Assistance API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package M5_NetworkAssistance

type IpPacketFilterSet struct {

	SrcIp string `json:"srcIp,omitempty"`

	DstIp string `json:"dstIp,omitempty"`

	Protocol int32 `json:"protocol,omitempty"`

	SrcPort int32 `json:"srcPort,omitempty"`

	DstPort int32 `json:"dstPort,omitempty"`

	ToSTc string `json:"toSTc,omitempty"`

	FlowLabel int32 `json:"flowLabel,omitempty"`

	Spi int32 `json:"spi,omitempty"`

	Direction string `json:"direction"`
}
