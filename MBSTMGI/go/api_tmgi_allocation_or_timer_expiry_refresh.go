/*
 * 3gpp-mbs-tmgi
 *
 * API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSTMGI

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllocateTmgi - Request the allocation of TMGI(s) for new MBS session(s) or the refresh of the expiry time of already allocated TMGI(s).
func AllocateTmgi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
