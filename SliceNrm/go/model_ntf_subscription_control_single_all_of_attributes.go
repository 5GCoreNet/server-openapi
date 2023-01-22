/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm

type NtfSubscriptionControlSingleAllOfAttributes struct {

	NotificationRecipientAddress string `json:"notificationRecipientAddress,omitempty"`

	NotificationTypes []NotificationType `json:"notificationTypes,omitempty"`

	Scope Scope `json:"scope,omitempty"`

	// The filter format shall be compliant to XPath 1.0.
	NotificationFilter string `json:"notificationFilter,omitempty"`
}
