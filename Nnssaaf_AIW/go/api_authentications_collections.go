/*
 * Nnssaaf_AIW
 *
 * AAA Interworking Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnssaaf_AIW

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAuthenticationContext - Create authentication context
func CreateAuthenticationContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}