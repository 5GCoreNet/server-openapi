/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_SDM

// UcSubscriptionData - Contains the User Consent Subscription Data.
type UcSubscriptionData struct {

	// A map(list of key-value pairs) where user consent purpose serves as key of user consent
	UserConsentPerPurposeList map[string]UserConsent `json:"userConsentPerPurposeList,omitempty"`
}
