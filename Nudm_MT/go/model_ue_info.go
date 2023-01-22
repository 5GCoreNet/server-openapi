/*
 * Nudm_MT
 *
 * UDM MT Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_MT

// UeInfo - Represents UE information.
type UeInfo struct {

	TadsInfo UeContextInfo `json:"tadsInfo,omitempty"`

	UserState Model5GsUserState `json:"userState,omitempty"`

	Var5gSrvccInfo Model5GSrvccInfo `json:"5gSrvccInfo,omitempty"`
}
