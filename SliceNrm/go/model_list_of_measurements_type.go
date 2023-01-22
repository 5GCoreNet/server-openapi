/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm

// ListOfMeasurementsType - See details in 3GPP TS 32.422 clause 5.10.3 for details.
type ListOfMeasurementsType struct {

	UMTS []string `json:"UMTS,omitempty"`

	LTE []string `json:"LTE,omitempty"`

	NR []string `json:"NR,omitempty"`
}
