/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadSessionManagementPolicyData - Retrieves the session management policy data for a subscriber
func ReadSessionManagementPolicyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateSessionManagementPolicyData - Modify the session management policy data for a subscriber
func UpdateSessionManagementPolicyData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
