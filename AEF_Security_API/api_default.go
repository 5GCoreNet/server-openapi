/*
 * AEF_Security_API
 *
 * API for AEF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AEF_Security_API

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckAuthenticationPost - Check authentication.
func CheckAuthenticationPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RevokeAuthorizationPost - Revoke authorization.
func RevokeAuthorizationPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
