/*
 * VAE_Service Continuity
 *
 * API for VAE Service Continuity Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_ServiceContinuity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryServiceContinuity - VAE service continuity query service operation
func QueryServiceContinuity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
