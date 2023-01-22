/*
 * Nnef_EASDeployment
 *
 * NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EASDeployment

// EasDeployInfoNotif - Represents notifications on EAS Deployment Information changes event(s) that occurred for an  Individual EAS Deployment Event Subscription resource. 
type EasDeployInfoNotif struct {

	EasDepNotifs []EasDepNotification `json:"easDepNotifs"`

	NotifId string `json:"notifId"`
}
