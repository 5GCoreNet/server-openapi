/*
 * Unified Data Repository Service API file for subscription data
 *
 * Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Subscription_Data

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ModifysdmSubscription - Modify an individual sdm subscription
func ModifysdmSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// QuerysdmSubscription - Retrieves a individual sdmSubscription identified by subsId
func QuerysdmSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RemovesdmSubscriptions - Deletes a sdmsubscriptions
func RemovesdmSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Updatesdmsubscriptions - Update an individual sdm subscriptions of a UE
func Updatesdmsubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}