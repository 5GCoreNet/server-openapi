/*
 * Unified Data Repository Service API file for structured data for exposure
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Exposure_Data

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
		"/",
		Index,
	},

	{
		"CreateOrReplaceAccessAndMobilityData",
		http.MethodPut,
		"/exposure-data/:ueId/access-and-mobility-data",
		CreateOrReplaceAccessAndMobilityData,
	},

	{
		"DeleteAccessAndMobilityData",
		http.MethodDelete,
		"/exposure-data/:ueId/access-and-mobility-data",
		DeleteAccessAndMobilityData,
	},

	{
		"QueryAccessAndMobilityData",
		http.MethodGet,
		"/exposure-data/:ueId/access-and-mobility-data",
		QueryAccessAndMobilityData,
	},

	{
		"UpdateAccessAndMobilityData",
		http.MethodPatch,
		"/exposure-data/:ueId/access-and-mobility-data",
		UpdateAccessAndMobilityData,
	},

	{
		"CreateIndividualExposureDataSubscription",
		http.MethodPost,
		"/exposure-data/subs-to-notify",
		CreateIndividualExposureDataSubscription,
	},

	{
		"DeleteIndividualExposureDataSubscription",
		http.MethodDelete,
		"/exposure-data/subs-to-notify/:subId",
		DeleteIndividualExposureDataSubscription,
	},

	{
		"ReplaceIndividualExposureDataSubscription",
		http.MethodPut,
		"/exposure-data/subs-to-notify/:subId",
		ReplaceIndividualExposureDataSubscription,
	},

	{
		"CreateOrReplaceSessionManagementData",
		http.MethodPut,
		"/exposure-data/:ueId/session-management-data/:pduSessionId",
		CreateOrReplaceSessionManagementData,
	},

	{
		"DeleteSessionManagementData",
		http.MethodDelete,
		"/exposure-data/:ueId/session-management-data/:pduSessionId",
		DeleteSessionManagementData,
	},

	{
		"QuerySessionManagementData",
		http.MethodGet,
		"/exposure-data/:ueId/session-management-data/:pduSessionId",
		QuerySessionManagementData,
	},
}
