/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudsf_DataRepository

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrModifyBlock - Create or Update a specific Block in a Record.
func CreateOrModifyBlock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteBlock - Delete a specific Block. Then update the Record
func DeleteBlock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetBlock - Retrieve a specific Block
func GetBlock(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetBlockList - Record's Blocks access
func GetBlockList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
