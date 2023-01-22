/*
 * Ndcaf_DataReporting
 *
 * Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndcaf_DataReporting

type CacheStatusAnyOf string

// List of CacheStatusAnyOf
const (
	HIT CacheStatusAnyOf = "HIT"
	MISS CacheStatusAnyOf = "MISS"
	EXPIRED CacheStatusAnyOf = "EXPIRED"
)
