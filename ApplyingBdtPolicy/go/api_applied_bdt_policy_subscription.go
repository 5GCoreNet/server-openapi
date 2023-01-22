/*
 * 3gpp-applying-bdt-policy
 *
 * API for applying BDT policy   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ApplyingBdtPolicy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNewSubscription - Creates a new subscription resource
func CreateNewSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ReadAllSubscriptions - read all of the active subscriptions for the AF
func ReadAllSubscriptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
