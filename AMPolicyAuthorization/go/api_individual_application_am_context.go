/*
 * 3gpp-am-policyauthorization
 *
 * API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AMPolicyAuthorization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAppAmContext - Deletes an existing Individual Application AM Context
func DeleteAppAmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAppAmContext - read an existing Individual application AM context
func GetAppAmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModAppAmContext - partial modifies an existing Individual application AM context
func ModAppAmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
