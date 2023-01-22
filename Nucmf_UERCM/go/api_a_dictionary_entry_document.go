/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nucmf_UERCM

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDictionaryEntry - Create a dictionary entry in the UCMF
func CreateDictionaryEntry(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
