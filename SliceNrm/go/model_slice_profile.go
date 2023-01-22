/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm

type SliceProfile struct {

	ServiceProfileId string `json:"serviceProfileId,omitempty"`

	PlmnInfoList []PlmnInfo `json:"plmnInfoList,omitempty"`

	CNSliceSubnetProfile CnSliceSubnetProfile `json:"cNSliceSubnetProfile,omitempty"`

	RANSliceSubnetProfile RanSliceSubnetProfile `json:"rANSliceSubnetProfile,omitempty"`

	TopSliceSubnetProfile TopSliceSubnetProfile `json:"topSliceSubnetProfile,omitempty"`
}
