/*
 * Nudm_UECM
 *
 * Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudm_UECM

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PeiUpdate - Updates the PEI in the 3GPP access registration context
func PeiUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}