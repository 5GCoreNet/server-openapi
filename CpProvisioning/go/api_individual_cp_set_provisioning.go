/*
 * 3gpp-cp-parameter-provisioning
 *
 * API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CpProvisioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndCPSetProvisioning - Delete CP at individual CP set(s) level associated with a CP parameter set Id.
func DeleteIndCPSetProvisioning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchIndCPSetProvisioning - Read CP at individual CP set(s) level associated with a CP parameter set Id.
func FetchIndCPSetProvisioning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateIndCPSetProvisioning - Update CP at individual CP set(s) level associated with a CP parameter set Id.
func UpdateIndCPSetProvisioning(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
