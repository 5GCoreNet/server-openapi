/*
 * 3gpp-ecs-address-provision
 *
 * API for ECS Address Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EcsAddressProvision

// GadShape - Common base type for GAD shapes.
type GadShape struct {

	Shape SupportedGadShapes `json:"shape"`
}