/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Application_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateIndividualInfluenceDataSubscription - Create a new Individual Influence Data Subscription resource
func CreateIndividualInfluenceDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadInfluenceDataSubscriptions - Read Influence Data Subscriptions
func ReadInfluenceDataSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
