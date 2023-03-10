/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateADRFDataStoreRecord - Creates a new Individual Data Store Record resource.
func CreateADRFDataStoreRecord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAdrfDataStoreRecords - Retrieves existing Individual ADRF Data Store Records.
func GetAdrfDataStoreRecords(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
