/*
 * 3gpp-monitoring-event
 *
 * API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MonitoringEvent

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteIndMonitoringEventSubscription - Deletes an already existing monitoring event subscription.
func DeleteIndMonitoringEventSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchIndMonitoringEventSubscription - Read an active subscriptions for the SCS/AS and the subscription Id.
func FetchIndMonitoringEventSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ModifyIndMonitoringEventSubscription - Modifies an existing subscription of monitoring event.
func ModifyIndMonitoringEventSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateIndMonitoringEventSubscription - Updates/replaces an existing subscription resource.
func UpdateIndMonitoringEventSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
