/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Ndccf_DataManagement

// ResourceUsage - The current usage of the virtual resources assigned to the NF instances belonging to a  particular network slice instance. 
type ResourceUsage struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	CpuUsage int32 `json:"cpuUsage,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MemoryUsage int32 `json:"memoryUsage,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	StorageUsage int32 `json:"storageUsage,omitempty"`
}
