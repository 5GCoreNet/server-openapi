/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_PP

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update - provision parameters
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
