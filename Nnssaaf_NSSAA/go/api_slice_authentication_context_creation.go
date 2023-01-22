/*
 * Nnssaaf_NSSAA
 *
 * Network Slice-Specific Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssaaf_NSSAA

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSliceAuthenticationContext - Create slice authentication context
func CreateSliceAuthenticationContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}