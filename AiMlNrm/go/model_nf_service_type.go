/*
 * AI/ML NRM
 *
 * OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AiMlNrm

type NfServiceType string

// List of NfServiceType
const (
	NAMF_COMMUNICATION NfServiceType = "Namf_Communication"
	NAMF_EVENT_EXPOSURE NfServiceType = "Namf_EventExposure"
	NAMF_MT NfServiceType = "Namf_MT"
	NAMF_LOCATION NfServiceType = "Namf_Location"
	NSMF_PDU_SESSION NfServiceType = "Nsmf_PDUSession"
	NSMF_EVENT_EXPOSURE NfServiceType = "Nsmf_EventExposure"
	OTHERS NfServiceType = "Others"
)
