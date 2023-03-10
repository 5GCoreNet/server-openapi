/*
 * VAE_ApplicationRequirement
 *
 * API for VAE Application Requirement Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_VAE_ApplicationRequirement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateApplicationRequirement - VAE_Application_Requirements resource create service Operation
func CreateApplicationRequirement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
