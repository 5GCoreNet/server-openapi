/*
 * Ngmlc_Location
 *
 * GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ngmlc_Location

// LcsServiceAuth - Possible values are: - \"LOCATION_ALLOWED_WITH_NOTIFICATION\": Location allowed with notification - \"LOCATION_ALLOWED_WITHOUT_NOTIFICATION\": Location allowed without notification - \"LOCATION_ALLOWED_WITHOUT_RESPONSE\": Location with notification and privacy    verification; location allowed if no response - \"LOCATION_RESTRICTED_WITHOUT_RESPONSE\": Location with notification and privacy   verification; location restricted if no response - \"NOTIFICATION_ONLY\": Notification only - \"NOTIFICATION_AND_VERIFICATION_ONLY\": Notification and privacy verification only 
type LcsServiceAuth struct {
}
