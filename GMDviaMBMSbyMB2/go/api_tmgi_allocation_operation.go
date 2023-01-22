/*
 * GMDviaMBMSbyMB2
 *
 * API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package GMDviaMBMSbyMB2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTMGIAllocation - Creates a new TMGI Allocation resource for a given SCS/AS.
func CreateTMGIAllocation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchAllTMGIAllocations - read all TMGI Allocation resource for a given SCS/AS
func FetchAllTMGIAllocations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
