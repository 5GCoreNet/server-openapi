/*
 * MSGS_MSGDelivery
 *
 * API for MSGG MSGin5G Server Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package MSGS_MSGDelivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeliverUeMessagePost - UE deliver message to MSGin5G Server
func DeliverUeMessagePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
