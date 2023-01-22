/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

type NrfInfoServedGmlcInfoValue struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	GmlcNumbers []string `json:"gmlcNumbers,omitempty"`
}
