/*
 * LMF Location
 *
 * LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nlmf_Location

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CancelLocation - request cancellation of periodic or triggered location
func CancelLocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
