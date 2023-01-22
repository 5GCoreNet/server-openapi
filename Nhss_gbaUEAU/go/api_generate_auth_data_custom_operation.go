/*
 * Nhss_gbaUEAU
 *
 * Nhss UE Authentication Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_gbaUEAU

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateAuthData - Generate GBA authentication data for the UE
func GenerateAuthData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
