/*
 * Namf_Location
 *
 * AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Location

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CancelLocation - Namf_Location CancelLocation service operation
func CancelLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ProvideLocationInfo - Namf_Location ProvideLocationInfo service Operation
func ProvideLocationInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ProvidePositioningInfo - Namf_Location ProvidePositioningInfo service Operation
func ProvidePositioningInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}