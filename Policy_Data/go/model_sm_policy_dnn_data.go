/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Policy_Data

// SmPolicyDnnData - Contains the SM policy data for a given DNN (and S-NSSAI).
type SmPolicyDnnData struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`

	AllowedServices []string `json:"allowedServices,omitempty"`

	SubscCats []string `json:"subscCats,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GbrUl string `json:"gbrUl,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GbrDl string `json:"gbrDl,omitempty"`

	AdcSupport bool `json:"adcSupport,omitempty"`

	SubscSpendingLimits bool `json:"subscSpendingLimits,omitempty"`

	// Represents information that identifies which IP pool or external server is used to allocate the IP address. 
	Ipv4Index int32 `json:"ipv4Index,omitempty"`

	// Represents information that identifies which IP pool or external server is used to allocate the IP address. 
	Ipv6Index int32 `json:"ipv6Index,omitempty"`

	Offline bool `json:"offline,omitempty"`

	Online bool `json:"online,omitempty"`

	ChfInfo ChargingInformation `json:"chfInfo,omitempty"`

	// A reference to the UsageMonitoringDataLimit or UsageMonitoringData instances for this DNN and SNSSAI that may also include the related monitoring key(s). The key of the map is the limit identifier. 
	RefUmDataLimitIds map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`

	MpsPriority bool `json:"mpsPriority,omitempty"`

	McsPriority bool `json:"mcsPriority,omitempty"`

	ImsSignallingPrio bool `json:"imsSignallingPrio,omitempty"`

	MpsPriorityLevel int32 `json:"mpsPriorityLevel,omitempty"`

	McsPriorityLevel int32 `json:"mcsPriorityLevel,omitempty"`

	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraInfos map[string]PresenceInfo `json:"praInfos,omitempty"`

	// Identifies transfer policies of background data transfer. Any string value can be used as a key of the map. 
	BdtRefIds *map[string]string `json:"bdtRefIds,omitempty"`

	LocRoutNotAllowed bool `json:"locRoutNotAllowed,omitempty"`
}
