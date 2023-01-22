/*
 * File Data Reporting MnS
 *
 * OAS 3.0.1 definition of the File Data Reporting MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package FileDataReportingMnS

type FileNotificationTypes string

// List of FileNotificationTypes
const (
	NOTIFY_FILE_READY FileNotificationTypes = "notifyFileReady"
	NOTIFY_FILE_PREPARATION_ERROR FileNotificationTypes = "notifyFilePreparationError"
)
