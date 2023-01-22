/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryGroupEEData - Retrieves the ee profile data profile data of a group or anyUE
func QueryGroupEEData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
