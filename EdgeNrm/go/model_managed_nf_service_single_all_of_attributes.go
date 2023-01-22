/*
 * 3GPP Edge NRM
 *
 * OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EdgeNrm

type ManagedNfServiceSingleAllOfAttributes struct {

	UserLabel string `json:"userLabel,omitempty"`

	NFServiceType NfServiceType `json:"nFServiceType,omitempty"`

	SAP Sap `json:"sAP,omitempty"`

	Operations []Operation `json:"operations,omitempty"`

	AdministrativeState AdministrativeState `json:"administrativeState,omitempty"`

	OperationalState OperationalState `json:"operationalState,omitempty"`

	UsageState UsageState `json:"usageState,omitempty"`

	RegistrationState RegistrationState `json:"registrationState,omitempty"`
}
