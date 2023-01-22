/*
 * 3gpp-nidd
 *
 * API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDD

// NiddConfigurationStatusNotification - Represents an NIDD configuration status notification.
type NiddConfigurationStatusNotification struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NiddConfiguration string `json:"niddConfiguration"`

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information.
	ExternalId string `json:"externalId,omitempty"`

	// string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN.
	Msisdn string `json:"msisdn,omitempty"`

	Status NiddStatus `json:"status"`

	// It indicates whether the network capability for the reliable data service is enabled or not.
	RdsCapIndication bool `json:"rdsCapIndication,omitempty"`

	RdsPort RdsPort `json:"rdsPort,omitempty"`
}
