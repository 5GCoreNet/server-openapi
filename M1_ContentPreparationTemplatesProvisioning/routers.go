/*
 * M1_ContentPreparationTemplatesProvisioning
 *
 * 5GMS AF M1 Content Preparation Templates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M1_ContentPreparationTemplatesProvisioning

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
		"CreateContentPreparationTemplate",
		http.MethodPost,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/content-preparation-templates",
		CreateContentPreparationTemplate,
	},

	{
		"DestroyContentPreparationTemplate",
		http.MethodDelete,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/content-preparation-templates/:contentPreparationTemplateId",
		DestroyContentPreparationTemplate,
	},

	{
		"PatchContentPreparationTemplate",
		http.MethodPatch,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/content-preparation-templates/:contentPreparationTemplateId",
		PatchContentPreparationTemplate,
	},

	{
		"RetrieveContentPreparationTemplate",
		http.MethodGet,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/content-preparation-templates/:contentPreparationTemplateId",
		RetrieveContentPreparationTemplate,
	},

	{
		"UpdateContentPreparationTemplate",
		http.MethodPut,
		"/3gpp-m1/v2/provisioning-sessions/:provisioningSessionId/content-preparation-templates/:contentPreparationTemplateId",
		UpdateContentPreparationTemplate,
	},
}
