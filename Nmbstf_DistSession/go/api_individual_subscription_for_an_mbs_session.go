/*
 * Nmbstf-distsession
 *
 * MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmbstf_DistSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusUnSubscribe - StatusUnSubscribe to unsubscribe from the Status Subscription
func StatusUnSubscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
