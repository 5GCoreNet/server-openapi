/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudsf_DataRepository

// RecordNotification - Definition of a notification on a record
type RecordNotification struct {

	Descriptor NotificationDescription `json:"descriptor"`

	Meta RecordMeta `json:"meta"`

	// list of opaque Block's in this Record
	Blocks []interface{} `json:"blocks,omitempty"`
}
