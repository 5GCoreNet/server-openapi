/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

type PolicyControlRequestTriggerAnyOf string

// List of PolicyControlRequestTriggerAnyOf
const (
	PLMN_CH PolicyControlRequestTriggerAnyOf = "PLMN_CH"
	RES_MO_RE PolicyControlRequestTriggerAnyOf = "RES_MO_RE"
	AC_TY_CH PolicyControlRequestTriggerAnyOf = "AC_TY_CH"
	UE_IP_CH PolicyControlRequestTriggerAnyOf = "UE_IP_CH"
	UE_MAC_CH PolicyControlRequestTriggerAnyOf = "UE_MAC_CH"
	AN_CH_COR PolicyControlRequestTriggerAnyOf = "AN_CH_COR"
	US_RE PolicyControlRequestTriggerAnyOf = "US_RE"
	APP_STA PolicyControlRequestTriggerAnyOf = "APP_STA"
	APP_STO PolicyControlRequestTriggerAnyOf = "APP_STO"
	AN_INFO PolicyControlRequestTriggerAnyOf = "AN_INFO"
	CM_SES_FAIL PolicyControlRequestTriggerAnyOf = "CM_SES_FAIL"
	PS_DA_OFF PolicyControlRequestTriggerAnyOf = "PS_DA_OFF"
	DEF_QOS_CH PolicyControlRequestTriggerAnyOf = "DEF_QOS_CH"
	SE_AMBR_CH PolicyControlRequestTriggerAnyOf = "SE_AMBR_CH"
	QOS_NOTIF PolicyControlRequestTriggerAnyOf = "QOS_NOTIF"
	NO_CREDIT PolicyControlRequestTriggerAnyOf = "NO_CREDIT"
	REALLO_OF_CREDIT PolicyControlRequestTriggerAnyOf = "REALLO_OF_CREDIT"
	PRA_CH PolicyControlRequestTriggerAnyOf = "PRA_CH"
	SAREA_CH PolicyControlRequestTriggerAnyOf = "SAREA_CH"
	SCNN_CH PolicyControlRequestTriggerAnyOf = "SCNN_CH"
	RE_TIMEOUT PolicyControlRequestTriggerAnyOf = "RE_TIMEOUT"
	RES_RELEASE PolicyControlRequestTriggerAnyOf = "RES_RELEASE"
	SUCC_RES_ALLO PolicyControlRequestTriggerAnyOf = "SUCC_RES_ALLO"
	RAI_CH PolicyControlRequestTriggerAnyOf = "RAI_CH"
	RAT_TY_CH PolicyControlRequestTriggerAnyOf = "RAT_TY_CH"
	REF_QOS_IND_CH PolicyControlRequestTriggerAnyOf = "REF_QOS_IND_CH"
	NUM_OF_PACKET_FILTER PolicyControlRequestTriggerAnyOf = "NUM_OF_PACKET_FILTER"
	UE_STATUS_RESUME PolicyControlRequestTriggerAnyOf = "UE_STATUS_RESUME"
	UE_TZ_CH PolicyControlRequestTriggerAnyOf = "UE_TZ_CH"
	AUTH_PROF_CH PolicyControlRequestTriggerAnyOf = "AUTH_PROF_CH"
	QOS_MONITORING PolicyControlRequestTriggerAnyOf = "QOS_MONITORING"
	SCELL_CH PolicyControlRequestTriggerAnyOf = "SCELL_CH"
	USER_LOCATION_CH PolicyControlRequestTriggerAnyOf = "USER_LOCATION_CH"
	EPS_FALLBACK PolicyControlRequestTriggerAnyOf = "EPS_FALLBACK"
	MA_PDU PolicyControlRequestTriggerAnyOf = "MA_PDU"
	TSN_BRIDGE_INFO PolicyControlRequestTriggerAnyOf = "TSN_BRIDGE_INFO"
	_5_G_RG_JOIN PolicyControlRequestTriggerAnyOf = "5G_RG_JOIN"
	_5_G_RG_LEAVE PolicyControlRequestTriggerAnyOf = "5G_RG_LEAVE"
	DDN_FAILURE PolicyControlRequestTriggerAnyOf = "DDN_FAILURE"
	DDN_DELIVERY_STATUS PolicyControlRequestTriggerAnyOf = "DDN_DELIVERY_STATUS"
	GROUP_ID_LIST_CHG PolicyControlRequestTriggerAnyOf = "GROUP_ID_LIST_CHG"
	DDN_FAILURE_CANCELLATION PolicyControlRequestTriggerAnyOf = "DDN_FAILURE_CANCELLATION"
	DDN_DELIVERY_STATUS_CANCELLATION PolicyControlRequestTriggerAnyOf = "DDN_DELIVERY_STATUS_CANCELLATION"
	VPLMN_QOS_CH PolicyControlRequestTriggerAnyOf = "VPLMN_QOS_CH"
	SUCC_QOS_UPDATE PolicyControlRequestTriggerAnyOf = "SUCC_QOS_UPDATE"
	SAT_CATEGORY_CHG PolicyControlRequestTriggerAnyOf = "SAT_CATEGORY_CHG"
	PCF_UE_NOTIF_IND PolicyControlRequestTriggerAnyOf = "PCF_UE_NOTIF_IND"
	NWDAF_DATA_CHG PolicyControlRequestTriggerAnyOf = "NWDAF_DATA_CHG"
)