/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

type SessionManagementSubscriptionData struct {

	SingleNssai Snssai `json:"singleNssai"`

	// A map (list of key-value pairs where Dnn, or optionally the Wildcard DNN, serves as key) of DnnConfigurations
	DnnConfigurations map[string]DnnConfiguration `json:"dnnConfigurations,omitempty"`

	InternalGroupIds []string `json:"internalGroupIds,omitempty"`

	// A map(list of key-value pairs) where GroupId serves as key of SharedDataId
	SharedVnGroupDataIds map[string]string `json:"sharedVnGroupDataIds,omitempty"`

	SharedDnnConfigurationsId string `json:"sharedDnnConfigurationsId,omitempty"`

	OdbPacketServices OdbPacketServices `json:"odbPacketServices,omitempty"`

	TraceData *TraceData `json:"traceData,omitempty"`

	SharedTraceDataId string `json:"sharedTraceDataId,omitempty"`

	// A map(list of key-value pairs) where Dnn serves as key of ExpectedUeBehaviourData
	ExpectedUeBehavioursList map[string]ExpectedUeBehaviourData `json:"expectedUeBehavioursList,omitempty"`

	// A map(list of key-value pairs) where Dnn serves as key of SuggestedPacketNumDl
	SuggestedPacketNumDlList map[string]SuggestedPacketNumDl `json:"suggestedPacketNumDlList,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
