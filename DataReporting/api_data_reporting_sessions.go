/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReporting

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDataRepSession - Create a new Data Reporting Session.
func CreateDataRepSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
