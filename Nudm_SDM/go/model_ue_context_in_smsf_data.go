/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

type UeContextInSmsfData struct {

	SmsfInfo3GppAccess SmsfInfo `json:"smsfInfo3GppAccess,omitempty"`

	SmsfInfoNon3GppAccess SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty"`
}
