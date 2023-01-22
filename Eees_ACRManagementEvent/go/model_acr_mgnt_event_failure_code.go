/*
 * EES ACR Management Event_API
 *
 * API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Eees_ACRManagementEvent

// AcrMgntEventFailureCode - Possible values are: - 3GPP_UP_PATH_CHANGE_MON_NOT_AVAILABLE: Indicates that the ACR Management Event Subscription failed because user plane path management event notifications from the 3GPP network is NOT available. This value is only applicable for the \"UP_PATH_CHG\", \"ACR_MONITORING\" and \"ACR_FACILITATION\" events. - OTHER_REASONS: Indicates that the ACR Management Event Subscription failed for other reasons. This value is applicable for all events. 
type AcrMgntEventFailureCode struct {
}
