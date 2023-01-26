/*
 * CAPIF_Events_API
 *
 * API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Events_API

// CapifEvent - Possible values are: - SERVICE_API_AVAILABLE: Events related to the availability of service APIs after the service APIs are published. - SERVICE_API_UNAVAILABLE: Events related to the unavailability of service APIs after the service APIs are unpublished. - SERVICE_API_UPDATE: Events related to change in service API information. - API_INVOKER_ONBOARDED: Events related to API invoker onboarded to CAPIF. - API_INVOKER_OFFBOARDED: Events related to API invoker offboarded from CAPIF. - SERVICE_API_INVOCATION_SUCCESS: Events related to the successful invocation of service APIs. - SERVICE_API_INVOCATION_FAILURE: Events related to the failed invocation of service APIs. - ACCESS_CONTROL_POLICY_UPDATE: Events related to the update for the access control policy related to the service APIs. - ACCESS_CONTROL_POLICY_UNAVAILABLE: Events related to the unavailability of the access control policy related to the service APIs. - API_INVOKER_AUTHORIZATION_REVOKED: Events related to the revocation of the authorization of API invokers to access the service APIs. - API_INVOKER_UPDATED: Events related to API invoker profile updated to CAPIF. - API_TOPOLOGY_HIDING_CREATED: Events related to the creation or update of the API topology hiding information of the service APIs after the service APIs are published. - API_TOPOLOGY_HIDING_REVOKED: Events related to the revocation of the API topology hiding information of the service APIs after the service APIs are unpublished. 
type CapifEvent struct {
}