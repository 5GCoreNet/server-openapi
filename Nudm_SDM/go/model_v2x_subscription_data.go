/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

type V2xSubscriptionData struct {

	NrV2xServicesAuth NrV2xAuth `json:"nrV2xServicesAuth,omitempty"`

	LteV2xServicesAuth LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	NrUePc5Ambr string `json:"nrUePc5Ambr,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	LtePc5Ambr string `json:"ltePc5Ambr,omitempty"`
}
