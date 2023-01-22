/*
 * Nnef_PFDmanagement Service API
 *
 * Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_PFDmanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NnefPFDmanagementCreateSubscr - Subscribe the notification of PFD changes.
func NnefPFDmanagementCreateSubscr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}