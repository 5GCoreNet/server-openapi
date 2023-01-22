/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Policy_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrReplaceUEPolicySet - Create or modify the UE policy set data for a subscriber
func CreateOrReplaceUEPolicySet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadUEPolicySet - Retrieves the UE policy set data for a subscriber
func ReadUEPolicySet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateUEPolicySet - Modify the UE policy set data for a subscriber
func UpdateUEPolicySet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
