/*
 * VAE_V2VConfigRequirement
 *
 * API for VAE_V2VConfigRequirement   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package VAE_V2VConfigRequirement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteV2VConfiguration - VAE V2V Configuration resource delete service Operation
func DeleteV2VConfiguration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadV2VConfiguration - VAE V2V Configuration resource read service Operation
func ReadV2VConfiguration(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}