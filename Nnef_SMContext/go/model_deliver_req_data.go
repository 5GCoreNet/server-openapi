/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_SMContext

// DeliverReqData - The data for Deliver service request, including the Mobile Originated data to be delivered via NEF.
type DeliverReqData struct {

	Data RefToBinaryData `json:"data"`
}
