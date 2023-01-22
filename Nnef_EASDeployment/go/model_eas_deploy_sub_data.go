/*
 * Nnef_EASDeployment
 *
 * NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EASDeployment

// EasDeploySubData - Represents an Individual EAS Deployment Event Subscription resource.
type EasDeploySubData struct {

	AppId string `json:"appId,omitempty"`

	// Each of the element identifies a (DNN, S-NSSAI) combination.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`

	EventId EasEvent `json:"eventId"`

	// Represents the EAS Deployment Information changes event(s) to be reported. Shall only be present if the \"immRep\" attribute is included and sets to true, and the current status of EAS Deployment Information is available. 
	EventsNotifs []EasDeployInfoData `json:"eventsNotifs,omitempty"`

	// Indication of immediate reporting. Set to true: requires the immediate reporting of the  current status of EAS Deployment Information, if available. Set to false (default): EAS  Deployment Information event report occurs when the event is met. 
	ImmRep bool `json:"immRep,omitempty"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId string `json:"interGroupId,omitempty"`

	NotifId string `json:"notifId"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
}
