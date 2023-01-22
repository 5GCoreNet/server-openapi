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

// CreateIndividualApplicationDataSubscription - Create a subscription to receive notification of application data changes
func CreateIndividualApplicationDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadApplicationDataChangeSubscriptions - Read Application Data change Subscriptions
func ReadApplicationDataChangeSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
