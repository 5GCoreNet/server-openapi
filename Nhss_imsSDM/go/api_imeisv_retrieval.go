/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsSDM

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIMEISVInfo - Retrieve the IMEISV information
func GetIMEISVInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
