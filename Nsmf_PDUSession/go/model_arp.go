/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

// Arp - Contains Allocation and Retention Priority information.
type Arp struct {

	// nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.  
	PriorityLevel *int32 `json:"priorityLevel"`

	PreemptCap PreemptionCapability `json:"preemptCap"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}
