/*
 * 3gpp-data-reporting-provisioning
 *
 * API for 3GPP Data Reporting and Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReportingProvisioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/3gpp-data-reporting-provisioning/v1/",
		Index,
	},

	{
		"CreateDataRepConfig",
		http.MethodPost,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId/configurations",
		CreateDataRepConfig,
	},

	{
		"CreateDataRepProvSession",
		http.MethodPost,
		"/3gpp-data-reporting-provisioning/v1/sessions",
		CreateDataRepProvSession,
	},

	{
		"DeleteIndDataRepConfig",
		http.MethodDelete,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId/configurations/:configurationId",
		DeleteIndDataRepConfig,
	},

	{
		"GetIndDataRepConfig",
		http.MethodGet,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId/configurations/:configurationId",
		GetIndDataRepConfig,
	},

	{
		"ModifyIndDataRepConfig",
		http.MethodPatch,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId/configurations/:configurationId",
		ModifyIndDataRepConfig,
	},

	{
		"UpdateIndDataRepConfig",
		http.MethodPut,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId/configurations/:configurationId",
		UpdateIndDataRepConfig,
	},

	{
		"DeleteIndDataRepProvSession",
		http.MethodDelete,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId",
		DeleteIndDataRepProvSession,
	},

	{
		"GetIndDataRepProvSession",
		http.MethodGet,
		"/3gpp-data-reporting-provisioning/v1/sessions/:sessionId",
		GetIndDataRepProvSession,
	},
}
