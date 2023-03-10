/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_EdgeNrm

type AffinityAntiAffinity struct {

	AffinityEAS []string `json:"affinityEAS,omitempty"`

	AntiAffinityEAS string `json:"antiAffinityEAS,omitempty"`
}
