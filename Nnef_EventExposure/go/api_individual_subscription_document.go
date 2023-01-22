/*
 * Nnef_EventExposure
 *
 * NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnef_EventExposure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndividualSubcription - unsubscribe from notifications
func DeleteIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetIndividualSubcription - retrieve subscription
func GetIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReplaceIndividualSubcription - update subscription
func ReplaceIndividualSubcription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
