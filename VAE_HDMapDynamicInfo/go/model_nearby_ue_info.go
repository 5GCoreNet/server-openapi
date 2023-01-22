/*
 * VAE_HDMapDynamicInfo
 *
 * API for VAE HDMapDynamicInfo Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_HDMapDynamicInfo

// NearbyUeInfo - Represents the informaiotn of nearby UEs.
type NearbyUeInfo struct {

	// Represents the identifier of the V2X UE.
	NearbyUeId string `json:"nearbyUeId"`

	Location UserLocation `json:"location"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Distance int32 `json:"distance"`
}
