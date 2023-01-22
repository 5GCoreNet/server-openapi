/*
 * Nsmsf_SMService Service API
 *
 * SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmsf_SMService

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SMServiceActivation - Activate SMS Service for a given UE
func SMServiceActivation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
