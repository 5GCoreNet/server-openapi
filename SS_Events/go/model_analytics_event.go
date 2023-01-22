/*
 * SS_Events
 *
 * API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SS_Events

// AnalyticsEvent - Possible values are: - UE_MOBILITY: The AF requests to be notified about analytics information of UE mobility. - UE_COMM: The AF requests to be notified about analytics information of UE communication. - ABNORMAL_BEHAVIOR: The AF requests to be notified about analytics information of UE's abnormal behavior. - CONGESTION: The AF requests to be notified about analytics information of user data congestion information.  - NETWORK_PERFORMANCE: The AF requests to be notified about analytics information of network performance.  - QOS_SUSTAINABILITY: The AF requests to be notified about analytics information of QoS sustainability. - DISPERSION: The AF requests to be notified about analytics information of Dispersion analytics. - DN_PERFORMANCE: The AF requests to be notified about analytics information of DN performance. - SERVICE_EXPERIENCE: The AF requests to be notified about analytics information of service experience. 
type AnalyticsEvent struct {
}