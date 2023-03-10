/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusSubscribeMod - StatusSubscribe to modify (update or renew) an individual subscription
func StatusSubscribeMod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// StatusUnSubscribe - StatusUnSubscribe to unsubscribe from the Status Subscription
func StatusUnSubscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
