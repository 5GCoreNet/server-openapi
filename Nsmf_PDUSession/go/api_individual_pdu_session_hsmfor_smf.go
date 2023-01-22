/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nsmf_PDUSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReleasePduSession - Release
func ReleasePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RetrievePduSession - Retrieve
func RetrievePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// TransferMoData - Transfer MO Data
func TransferMoData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdatePduSession - Update (initiated by V-SMF or I-SMF)
func UpdatePduSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
