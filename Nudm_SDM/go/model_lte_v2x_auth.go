/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

// LteV2xAuth - Contains LTE V2X services authorized information.
type LteV2xAuth struct {

	VehicleUeAuth UeAuth `json:"vehicleUeAuth,omitempty"`

	PedestrianUeAuth UeAuth `json:"pedestrianUeAuth,omitempty"`
}
