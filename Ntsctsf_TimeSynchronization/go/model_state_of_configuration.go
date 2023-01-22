/*
 * Ntsctsf_TimeSynchronization Service API
 *
 * TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ntsctsf_TimeSynchronization

// StateOfConfiguration - Contains the state of the time synchronization configuration.
type StateOfConfiguration struct {

	// When the PTP port state is Leader, Follower or Passive, it is included and set to true to indicate the state of configuration for NW-TT port is active; when PTP port state is in any other case, it is included and set to false to indicate the state of configuration for NW-TT port is inactive. Default value is false. 
	StateNwtt bool `json:"stateNwtt,omitempty"`

	// Contains the PTP port states of the DS-TT(s).
	StateOfDstts []StateOfDstt `json:"stateOfDstts,omitempty"`
}
