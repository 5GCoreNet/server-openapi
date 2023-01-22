/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nbsf_Management

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndividualSubcription - Delete an individual subscription for event notifications from the BSF
func DeleteIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReplaceIndividualSubcription - Replace an individual subscription for event notifications from the BSF
func ReplaceIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
