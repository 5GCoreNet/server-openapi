/*
 * TS 28.532 Streaming data reporting service
 *
 * OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 17.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package StreamingDataMnS

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConnectionsConnectionIdGet - Obtain information about a connection.
func ConnectionsConnectionIdGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsConnectionIdStreamsDelete - Remove reporting streams from an existing connection
func ConnectionsConnectionIdStreamsDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsConnectionIdStreamsGet - Obtain information about streams.
func ConnectionsConnectionIdStreamsGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsConnectionIdStreamsPost - Inform consumer about new reporting streams on an existing connection.
func ConnectionsConnectionIdStreamsPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsConnectionIdStreamsStreamIdGet - Obtain information about stream
func ConnectionsConnectionIdStreamsStreamIdGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsGet - Obtain information about connections.
func ConnectionsGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConnectionsPost - Inform consumer about reporting streams to be carried by the new connection and receive a new connection id.
func ConnectionsPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}