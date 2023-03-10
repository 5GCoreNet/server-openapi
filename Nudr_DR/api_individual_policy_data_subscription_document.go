/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndividualPolicyDataSubscription - Delete the individual Policy Data subscription
func DeleteIndividualPolicyDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReplaceIndividualPolicyDataSubscription - Modify a subscription to receive notification of policy data changes
func ReplaceIndividualPolicyDataSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
