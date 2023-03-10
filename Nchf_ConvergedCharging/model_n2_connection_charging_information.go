/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

type N2ConnectionChargingInformation struct {

	N2ConnectionMessageType int32 `json:"n2ConnectionMessageType"`

	UserInformation UserInformation `json:"userInformation,omitempty"`

	UserLocationinfo UserLocation `json:"userLocationinfo,omitempty"`

	PSCellInformation PsCellInformation `json:"pSCellInformation,omitempty"`

	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	UetimeZone string `json:"uetimeZone,omitempty"`

	RATType RatType `json:"rATType,omitempty"`

	AmfUeNgapId int32 `json:"amfUeNgapId,omitempty"`

	RanUeNgapId int32 `json:"ranUeNgapId,omitempty"`

	RanNodeId GlobalRanNodeId `json:"ranNodeId,omitempty"`

	RestrictedRatList []RatType `json:"restrictedRatList,omitempty"`

	ForbiddenAreaList []Area `json:"forbiddenAreaList,omitempty"`

	ServiceAreaRestriction []ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	RestrictedCnList []CoreNetworkType `json:"restrictedCnList,omitempty"`

	AllowedNSSAI []Snssai `json:"allowedNSSAI,omitempty"`

	RrcEstCause string `json:"rrcEstCause,omitempty"`
}
