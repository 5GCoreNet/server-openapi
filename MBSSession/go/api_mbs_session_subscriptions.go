/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMBSSessionsSubsc - Request the creation of a new Individual MBS Session subscription resource.
func CreateMBSSessionsSubsc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadMBSSessionsSubscs - Retrieve all the active MBS Sessions subscriptions.
func ReadMBSSessionsSubscs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
