/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// NetworkPerfType - Possible values are: - GNB_ACTIVE_RATIO: Indicates that the network performance requirement is gNodeB active (i.e. up and running) rate. Indicates the ratio of gNB active (i.e. up and running) number to the total number of gNB - GNB_COMPUTING_USAGE: Indicates gNodeB computing resource usage. - GNB_MEMORY_USAGE: Indicates gNodeB memory usage. - GNB_DISK_USAGE: Indicates gNodeB disk usage. - NUM_OF_UE: Indicates number of UEs. - SESS_SUCC_RATIO: Indicates ratio of successful setup of PDU sessions to total PDU session setup attempts. - HO_SUCC_RATIO: Indicates Ratio of successful handovers to the total handover attempts.  
type NetworkPerfType struct {
}
