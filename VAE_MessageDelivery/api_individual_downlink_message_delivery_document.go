/*
 * VAE_MessageDelivery
 *
 * API for VAE Message Delivery Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_VAE_MessageDelivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReadIndividualDownlinkMessageDelivery - VAE Message delivery resource Read service Operation
func ReadIndividualDownlinkMessageDelivery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}