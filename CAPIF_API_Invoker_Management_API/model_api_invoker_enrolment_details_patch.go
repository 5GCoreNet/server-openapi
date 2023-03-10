/*
 * CAPIF_API_Invoker_Management_API
 *
 * API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_API_Invoker_Management_API

// ApiInvokerEnrolmentDetailsPatch - Represents an API Invoker's enrolment details to be updated.
type ApiInvokerEnrolmentDetailsPatch struct {

	OnboardingInformation OnboardingInformation `json:"onboardingInformation,omitempty"`

	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination string `json:"notificationDestination,omitempty"`

	// The list of service APIs that the API Invoker is allowed to invoke
	ApiList []ServiceApiDescription `json:"apiList,omitempty"`

	// Generic information related to the API invoker such as details of the device or the application. 
	ApiInvokerInformation string `json:"apiInvokerInformation,omitempty"`
}
