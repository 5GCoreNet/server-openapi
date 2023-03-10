/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndMBSSessionsSubsc - Request the deletion of an existing Individual MBS Session subscription resource.
func DeleteIndMBSSessionsSubsc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadIndMBSSessionsSubsc - Retrieve an existing Individual MBS Session Subscription resource.
func ReadIndMBSSessionsSubsc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
