/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MBSUserDataIngestSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMBSUserDataIngestSession - Request the creation of a new Individual MBS User Data Ingest Session resource.
func CreateMBSUserDataIngestSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RetrieveMBSUserDataIngestSessions - Retrieve all the active MBS User Data Ingest Sessions managed by the NEF.
func RetrieveMBSUserDataIngestSessions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}