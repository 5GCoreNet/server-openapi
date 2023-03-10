/*
 * M1_EdgeResourcesProvisioning
 *
 * 5GMS AF M1 Edge Resources Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_EdgeResourcesProvisioning

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
		"/3gpp-m1/v2/",
		Index,
	},

	{
		"CreateEdgeResourcesConfiguration",
		http.MethodPost,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/edge-resources-configurations",
		CreateEdgeResourcesConfiguration,
	},

	{
		"DestroyEdgeResourcesConfiguration",
		http.MethodDelete,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/edge-resources-configurations/:edgeResourcesConfigurationId",
		DestroyEdgeResourcesConfiguration,
	},

	{
		"PatchEdgeResourcesConfiguration",
		http.MethodPatch,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/edge-resources-configurations/:edgeResourcesConfigurationId",
		PatchEdgeResourcesConfiguration,
	},

	{
		"RetrieveEdgeResourcesConfiguration",
		http.MethodGet,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/edge-resources-configurations/:edgeResourcesConfigurationId",
		RetrieveEdgeResourcesConfiguration,
	},

	{
		"UpdateEdgeResourcesConfiguration",
		http.MethodPut,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/edge-resources-configurations/:edgeResourcesConfigurationId",
		UpdateEdgeResourcesConfiguration,
	},
}
