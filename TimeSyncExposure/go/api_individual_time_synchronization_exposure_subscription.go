/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TimeSyncExposure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAnSubscription - Deletes an already existing subscription
func DeleteAnSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FullyUpdateAnSubscription - Fully updates/replaces an existing subscription resource
func FullyUpdateAnSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadAnSubscription - read an active subscription for the AF and the subscription Id
func ReadAnSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadTimeSynSubscription - read an active subscription for the AF and the subscription Id
func ReadTimeSynSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}