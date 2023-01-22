/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsSDM

// ChargingInfo - Diameter addresses of the charging function
type ChargingInfo struct {

	// Fully Qualified Domain Name
	PrimaryEventChargingFunctionName string `json:"primaryEventChargingFunctionName,omitempty"`

	// Fully Qualified Domain Name
	SecondaryEventChargingFunctionName string `json:"secondaryEventChargingFunctionName,omitempty"`

	// Fully Qualified Domain Name
	PrimaryChargingCollectionFunctionName string `json:"primaryChargingCollectionFunctionName,omitempty"`

	// Fully Qualified Domain Name
	SecondaryChargingCollectionFunctionName string `json:"secondaryChargingCollectionFunctionName,omitempty"`
}
