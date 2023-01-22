/*
 * Slice NRM
 *
 * OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SliceNrm

// TraceReferenceType - The Trace Reference parameter shall be globally unique, therefore the Trace Reference shall compose as follows - MCC+MNC+Trace ID, where the MCC and MNC are coming with the Trace activation request from the management system to identify one PLMN containing the management system, and Trace ID is a 3 byte Octet String. See 3GPP TS 32.422 clause 5.6 for additional details.
type TraceReferenceType struct {

	Mcc string `json:"mcc"`

	Mnc string `json:"mnc"`

	TraceId string `json:"traceId"`
}
