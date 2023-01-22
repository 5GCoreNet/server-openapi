/*
 * File Data Reporting MnS
 *
 * OAS 3.0.1 definition of the File Data Reporting MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package FileDataReportingMnS

type CmNotificationTypes string

// List of CmNotificationTypes
const (
	NOTIFY_MOI_CREATION CmNotificationTypes = "notifyMOICreation"
	NOTIFY_MOI_DELETION CmNotificationTypes = "notifyMOIDeletion"
	NOTIFY_MOI_ATTRIBUTE_VALUE_CHANGES CmNotificationTypes = "notifyMOIAttributeValueChanges"
	NOTIFY_MOI_CHANGES CmNotificationTypes = "notifyMOIChanges"
)
