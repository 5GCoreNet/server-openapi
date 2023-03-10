/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type SmServiceTypeAnyOf string

// List of SmServiceTypeAnyOf
const (
	CONTENT_PROCESSING SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_CONTENT_PROCESSING"
	FORWARDING SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_FORWARDING"
	FORWARDING_MULTIPLE_SUBSCRIPTIONS SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_FORWARDING_MULTIPLE_SUBSCRIPTIONS"
	FILTERING SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_FILTERING"
	RECEIPT SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_RECEIPT"
	NETWORK_STORAGE SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_NETWORK_STORAGE"
	TO_MULTIPLE_DESTINATIONS SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_TO_MULTIPLE_DESTINATIONS"
	VIRTUAL_PRIVATE_NETWORK_VPN SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_VIRTUAL_PRIVATE_NETWORK(VPN)"
	AUTO_REPLY SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_AUTO_REPLY"
	PERSONAL_SIGNATURE SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_PERSONAL_SIGNATURE"
	DEFERRED_DELIVERY SmServiceTypeAnyOf = "VAS4SMS_SHORT_MESSAGE_DEFERRED_DELIVERY"
)
