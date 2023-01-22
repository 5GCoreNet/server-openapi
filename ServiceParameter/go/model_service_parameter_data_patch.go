/*
 * 3gpp-service-parameter
 *
 * API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ServiceParameter

// ServiceParameterDataPatch - Represents the parameters to request the modification of a service parameter subscription resource. 
type ServiceParameterDataPatch struct {

	// Represents the same as the ParameterOverPc5 data type but with the nullable:true property. 
	ParamOverPc5 *string `json:"paramOverPc5,omitempty"`

	// Represents the same as the ParameterOverUu data type but with the nullable:true property. 
	ParamOverUu *string `json:"paramOverUu,omitempty"`

	// This data type is defined in the same way as the ParamForProSeDd data type, but with the OpenAPI nullable property set to true. 
	ParamForProSeDd *string `json:"paramForProSeDd,omitempty"`

	// This data type is defined in the same way as the ParamForProSeDc data type, but with the OpenAPI nullable property set to true. 
	ParamForProSeDc *string `json:"paramForProSeDc,omitempty"`

	// This data type is defined in the same way as the ParamForProSeU2NRelay data type, but with the OpenAPI nullable property set to true. 
	ParamForProSeU2NRelUe *string `json:"paramForProSeU2NRelUe,omitempty"`

	// This data type is defined in the same way as the ParamForProSeRemUe data type, but with the OpenAPI nullable property set to true. 
	ParamForProSeRemUe *string `json:"paramForProSeRemUe,omitempty"`

	// Contains the service parameter used to guide the URSP.
	UrspGuidance []UrspRuleRequest `json:"urspGuidance,omitempty"`

	SubNotifEvents *[]Event `json:"subNotifEvents,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	NotificationDestination string `json:"notificationDestination,omitempty"`
}
