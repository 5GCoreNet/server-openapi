/*
 * 3gpp-nidd-configuration-trigger
 *
 * API for NIDD Configuration Trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDDConfigurationTrigger

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NiddConfigurationTrigger - 
func NiddConfigurationTrigger(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
