/*
 * Neasdf_BaselineDNSPattern
 *
 * EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Neasdf_BaselineDNSPattern

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrReplaceBaseDnsPattern - Creates or Updates the Baseline DNS Pattern (complete replacement)
func CreateOrReplaceBaseDnsPattern(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteBaseDnsPattern - Deletes a Baseline DNS Pattern
func DeleteBaseDnsPattern(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateBaseDNSPattern - Updates the Baseline DNS Pattern
func UpdateBaseDNSPattern(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
