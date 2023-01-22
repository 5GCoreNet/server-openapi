/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnrf_NFManagement

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
		"/nnrf-nfm/v1/",
		Index,
	},

	{
		"DeregisterNFInstance",
		http.MethodDelete,
		"/nnrf-nfm/v1/nf-instances/:nfInstanceID",
		DeregisterNFInstance,
	},

	{
		"GetNFInstance",
		http.MethodGet,
		"/nnrf-nfm/v1/nf-instances/:nfInstanceID",
		GetNFInstance,
	},

	{
		"RegisterNFInstance",
		http.MethodPut,
		"/nnrf-nfm/v1/nf-instances/:nfInstanceID",
		RegisterNFInstance,
	},

	{
		"UpdateNFInstance",
		http.MethodPatch,
		"/nnrf-nfm/v1/nf-instances/:nfInstanceID",
		UpdateNFInstance,
	},

	{
		"GetNFInstances",
		http.MethodGet,
		"/nnrf-nfm/v1/nf-instances",
		GetNFInstances,
	},

	{
		"OptionsNFInstances",
		http.MethodOptions,
		"/nnrf-nfm/v1/nf-instances",
		OptionsNFInstances,
	},

	{
		"RemoveSubscription",
		http.MethodDelete,
		"/nnrf-nfm/v1/subscriptions/:subscriptionID",
		RemoveSubscription,
	},

	{
		"UpdateSubscription",
		http.MethodPatch,
		"/nnrf-nfm/v1/subscriptions/:subscriptionID",
		UpdateSubscription,
	},

	{
		"CreateSubscription",
		http.MethodPost,
		"/nnrf-nfm/v1/subscriptions",
		CreateSubscription,
	},
}
