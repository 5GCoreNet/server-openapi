/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Subscription_Data

// DlDataDeliveryStatus - Possible values are: - BUFFERED: The first downlink data is buffered with extended buffering matching the   source of the downlink traffic. - TRANSMITTED: The first downlink data matching the source of the downlink traffic is   transmitted after previous buffering or discarding of corresponding packet(s) because   the UE of the PDU Session becomes ACTIVE, and buffered data can be delivered to UE. - DISCARDED: The first downlink data matching the source of the downlink traffic is   discarded because the Extended Buffering time, as determined by the SMF, expires or   the amount of downlink data to be buffered is exceeded. 
type DlDataDeliveryStatus struct {
}
