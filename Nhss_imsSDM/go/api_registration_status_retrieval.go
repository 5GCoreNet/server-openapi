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

// GetRegistrationStatus - Retrieve the registration status of a user
func GetRegistrationStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
