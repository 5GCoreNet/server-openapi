/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type LocationReportingChargingInformation struct {

	LocationReportingMessageType int32 `json:"locationReportingMessageType"`

	UserInformation UserInformation `json:"userInformation,omitempty"`

	UserLocationinfo UserLocation `json:"userLocationinfo,omitempty"`

	PSCellInformation PsCellInformation `json:"pSCellInformation,omitempty"`

	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone string `json:"uetimeZone,omitempty"`

	RATType RatType `json:"rATType,omitempty"`

	PresenceReportingAreaInformation map[string]PresenceInfo `json:"presenceReportingAreaInformation,omitempty"`
}
