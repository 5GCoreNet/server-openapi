/*
 * Nnsacf_NSAC
 *
 * Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnsacf_NSAC

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NumOfPDUsUpdate - Network Slice Admission Control on the number of PDU Sessions
func NumOfPDUsUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// NumOfUEsUpdate - Network Slice Admission Control on the Number of UEs
func NumOfUEsUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
