/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

// NwdafEvent - Possible values are: - SLICE_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice   - NETWORK_PERFORMANCE: Indicates that the event subscribed is network performance information.   - NF_LOAD: Indicates that the event subscribed is load level and status of one or several Network Functions.   - SERVICE_EXPERIENCE: Indicates that the event subscribed is service experience.   - UE_MOBILITY: Indicates that the event subscribed is UE mobility information.   - UE_COMMUNICATION: Indicates that the event subscribed is UE communication information.   - QOS_SUSTAINABILITY: Indicates that the event subscribed is QoS sustainability.   - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour.   - USER_DATA_CONGESTION: Indicates that the event subscribed is user data congestion information.   - NSI_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice and the optionally associated Network Slice Instance   - DN_PERFORMANCE: Indicates that the event subscribed is DN performance information.   - DISPERSION: Indicates that the event subscribed is dispersion information.   - RED_TRANS_EXP: Indicates that the event subscribed is redundant transmission experience.   - WLAN_PERFORMANCE: Indicates that the event subscribed is WLAN performance.   - SM_CONGESTION: Indicates the Session Management Congestion Control Experience information for specific DNN and/or S-NSSAI. 
type NwdafEvent struct {
}
