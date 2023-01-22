/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nadrf_DataManagement

// AfEventExposureNotif - Represents notifications on application event(s) that occurred for an Individual Application Event Subscription resource. 
type AfEventExposureNotif struct {

	NotifId string `json:"notifId"`

	EventNotifs []AfEventNotification `json:"eventNotifs"`
}
