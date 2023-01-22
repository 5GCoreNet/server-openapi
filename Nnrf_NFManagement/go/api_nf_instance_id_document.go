/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeregisterNFInstance - Deregisters a given NF Instance
func DeregisterNFInstance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetNFInstance - Read the profile of a given NF Instance
func GetNFInstance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RegisterNFInstance - Register a new NF Instance
func RegisterNFInstance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateNFInstance - Update NF Instance profile
func UpdateNFInstance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
