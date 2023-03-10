/*
 * Nhss_SDM
 *
 * HSS Subscriber Data Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nhss_SDM

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUeCtxInPgwData - Retrieve the UE Context In PGW
func GetUeCtxInPgwData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
