/*
 * Nhss_UECM
 *
 * HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_UECM

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeregisterSN - MME/SGSN Deregistration
func DeregisterSN(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
