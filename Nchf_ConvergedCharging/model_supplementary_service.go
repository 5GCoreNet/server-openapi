/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nchf_ConvergedCharging

import (
	"time"
)

type SupplementaryService struct {

	SupplementaryServiceType SupplementaryServiceType `json:"supplementaryServiceType,omitempty"`

	SupplementaryServiceMode SupplementaryServiceMode `json:"supplementaryServiceMode,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	NumberOfDiversions int32 `json:"numberOfDiversions,omitempty"`

	AssociatedPartyAddress string `json:"associatedPartyAddress,omitempty"`

	ConferenceId string `json:"conferenceId,omitempty"`

	ParticipantActionType ParticipantActionType `json:"participantActionType,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime time.Time `json:"changeTime,omitempty"`

	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	NumberOfParticipants int32 `json:"numberOfParticipants,omitempty"`

	CUGInformation string `json:"cUGInformation,omitempty"`
}
