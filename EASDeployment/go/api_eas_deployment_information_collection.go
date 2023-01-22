/*
 * 3gpp-eas-deployment
 *
 * API for AF provisioned EAS Deployment.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EASDeployment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAnDeployment - Create a new Individual EAS Deployment information resource.
func CreateAnDeployment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadAllDeployment - Read all EAS Deployment information for a given AF
func ReadAllDeployment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
