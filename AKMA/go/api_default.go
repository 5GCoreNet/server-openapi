/*
 * 3gpp-akma
 *
 * API for AKMA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package AKMA

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RetrieveAKMAAppKey - Retrieve AKMA Application Key Information.
func RetrieveAKMAAppKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
