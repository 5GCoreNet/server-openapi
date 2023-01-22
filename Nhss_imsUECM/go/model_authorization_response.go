/*
 * Nhss_imsUECM
 *
 * Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsUECM

// AuthorizationResponse - Ims Registration authorization information result
type AuthorizationResponse struct {

	AuthorizationResult AuthorizationResult `json:"authorizationResult"`

	CscfServerName string `json:"cscfServerName,omitempty"`

	ScscfSelectionAssistanceInfo ScscfSelectionAssistanceInformation `json:"scscfSelectionAssistanceInfo,omitempty"`
}
