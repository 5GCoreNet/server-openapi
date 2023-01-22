/*
 * N5g-ddnmf_Discovery API
 *
 * N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package N5g-ddnmf_Discovery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObtainMonitorAuth - Obtain the authorization to monitor for a UE
func ObtainMonitorAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
