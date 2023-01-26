/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_PfdManagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndPFDManagementTransaction - Delete PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).
func DeleteIndPFDManagementTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchIndPFDManagementTransaction - Read all PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).
func FetchIndPFDManagementTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifyIndPFDManagementTransaction - Modify an existing PFD Management Transaction resource.
func ModifyIndPFDManagementTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateIndPFDManagementTransaction - Update PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).
func UpdateIndPFDManagementTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}