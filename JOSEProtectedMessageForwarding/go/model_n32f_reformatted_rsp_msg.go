/*
 * JOSE Protected Message Forwarding API
 *
 * N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package JOSEProtectedMessageForwarding

// N32fReformattedRspMsg - Contains the reformatted HTTP/2 response message
type N32fReformattedRspMsg struct {

	ReformattedData FlatJweJson `json:"reformattedData"`

	ModificationsBlock []FlatJwsJson `json:"modificationsBlock,omitempty"`
}
