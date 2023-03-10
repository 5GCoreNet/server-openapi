/*
 * Nhss_imsUEAU
 *
 * Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_imsUEAU

// Model3GAkaAv - Contains RAND, XRES, AUTN, CK, and IK
type Model3GAkaAv struct {

	Rand string `json:"rand"`

	Xres string `json:"xres"`

	Autn string `json:"autn"`

	Ck string `json:"ck"`

	Ik string `json:"ik"`
}
