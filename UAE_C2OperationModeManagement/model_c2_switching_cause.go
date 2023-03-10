/*
 * UAE Server C2 Operation Mode Management Service
 *
 * UAE Server C2 Operation Mode Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_UAE_C2OperationModeManagement

// C2SwitchingCause - Possible values are: - DIRECT_LINK_QUALITY_DEGRADATION: Indicates that the C2 Communication Mode switching was triggered due to a degradation in the direct radio link quality. - DIRECT_LINK_AVAILABLE: Indicates that the C2 Communication Mode switching was triggered due to the availability of a direct link, i.e. direct radio link quality enables its usage. - MOVING_BVLOS: Indicates that the C2 Communication Mode switching was triggered due to the UAV moving BVLOS. - LOCATION_CHANGE: Indicates that the C2 Communication Mode switching was triggered due to an actual or expected location/mobility of the UAV (e.g. which impacts the UAV-to-UAV-C location). - TRAFFIC_CONTROL_NEEDED: Indicates that the C2 Communication Mode switching was triggered due to the necessity to have air traffic control. - SECURITY_REASONS: Indicates that the C2 Communication Mode switching was triggered due to security reasons. - OTHER_REASONS: Indicates that the C2 Communication Mode switching was triggered due to other reasons (e.g. unpredictable event, unknown reason, weather conditions, topography, etc.). 
type C2SwitchingCause struct {
}
