/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

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
		"/3gpp-mbs-ud-ingest/v1/",
		Index,
	},

	{
		"DeleteIndivMBSUserDataIngestSession",
		http.MethodDelete,
		"/3gpp-mbs-ud-ingest/v1/sessions/:sessionId",
		DeleteIndivMBSUserDataIngestSession,
	},

	{
		"ModifyIndivMBSUserDataIngestSession",
		http.MethodPatch,
		"/3gpp-mbs-ud-ingest/v1/sessions/:sessionId",
		ModifyIndivMBSUserDataIngestSession,
	},

	{
		"RetrieveIndivMBSUserDataIngestSession",
		http.MethodGet,
		"/3gpp-mbs-ud-ingest/v1/sessions/:sessionId",
		RetrieveIndivMBSUserDataIngestSession,
	},

	{
		"UpdateIndivMBSUserDataIngestSession",
		http.MethodPut,
		"/3gpp-mbs-ud-ingest/v1/sessions/:sessionId",
		UpdateIndivMBSUserDataIngestSession,
	},

	{
		"DeleteIndMBSUserDataIngStatSubsc",
		http.MethodDelete,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions/:subscriptionId",
		DeleteIndMBSUserDataIngStatSubsc,
	},

	{
		"ModifyIndMBSUserDataIngStatSubsc",
		http.MethodPatch,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions/:subscriptionId",
		ModifyIndMBSUserDataIngStatSubsc,
	},

	{
		"RetrieveIndMBSUserDataIngStatSubsc",
		http.MethodGet,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions/:subscriptionId",
		RetrieveIndMBSUserDataIngStatSubsc,
	},

	{
		"UpdateIndMBSUserDataIngStatSubsc",
		http.MethodPut,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions/:subscriptionId",
		UpdateIndMBSUserDataIngStatSubsc,
	},

	{
		"CreateMBSUserDataIngStatSubsc",
		http.MethodPost,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions",
		CreateMBSUserDataIngStatSubsc,
	},

	{
		"RetrieveMBSUserDataIngStatSubscs",
		http.MethodGet,
		"/3gpp-mbs-ud-ingest/v1/status-subscriptions",
		RetrieveMBSUserDataIngStatSubscs,
	},

	{
		"CreateMBSUserDataIngestSession",
		http.MethodPost,
		"/3gpp-mbs-ud-ingest/v1/sessions",
		CreateMBSUserDataIngestSession,
	},

	{
		"RetrieveMBSUserDataIngestSessions",
		http.MethodGet,
		"/3gpp-mbs-ud-ingest/v1/sessions",
		RetrieveMBSUserDataIngestSessions,
	},
}
