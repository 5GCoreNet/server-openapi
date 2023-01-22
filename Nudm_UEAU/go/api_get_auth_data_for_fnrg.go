/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_UEAU

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRgAuthData - Get authentication data for the FN-RG
func GetRgAuthData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
