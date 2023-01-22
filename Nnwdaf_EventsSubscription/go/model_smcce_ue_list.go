/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_EventsSubscription

// SmcceUeList - Represents the List of UEs classified based on experience level of Session Management  congestion control. 
type SmcceUeList struct {

	HighLevel []string `json:"highLevel,omitempty"`

	MediumLevel []string `json:"mediumLevel,omitempty"`

	LowLevel []string `json:"lowLevel,omitempty"`
}
