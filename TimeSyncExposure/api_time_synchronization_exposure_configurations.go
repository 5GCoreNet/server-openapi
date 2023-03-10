/*
 * 3gpp-time-sync-exposure
 *
 * API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_TimeSyncExposure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNewConfirguation - Creates a new configuration resource
func CreateNewConfirguation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadAllConfirguations - read all of the active configurations for the AF
func ReadAllConfirguations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
