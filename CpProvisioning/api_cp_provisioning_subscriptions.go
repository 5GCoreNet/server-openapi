/*
 * 3gpp-cp-parameter-provisioning
 *
 * API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CpProvisioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCPProvisioningSubscription - Create a new subscription resource of provisioning CP parameter set(s).
func CreateCPProvisioningSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchAllCPProvisioningSubscriptions - Read all active CP parameter provisioning subscription resources for a given SCS/AS.
func FetchAllCPProvisioningSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}