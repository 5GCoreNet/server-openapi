/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

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
		"ReadAccessAndMobilityPolicyData",
		http.MethodGet,
		"/policy-data/ues/:ueId/am-data",
		ReadAccessAndMobilityPolicyData,
	},

	{
		"ReadBdtData",
		http.MethodGet,
		"/policy-data/bdt-data",
		ReadBdtData,
	},

	{
		"CreateIndividualBdtData",
		http.MethodPut,
		"/policy-data/bdt-data/:bdtReferenceId",
		CreateIndividualBdtData,
	},

	{
		"DeleteIndividualBdtData",
		http.MethodDelete,
		"/policy-data/bdt-data/:bdtReferenceId",
		DeleteIndividualBdtData,
	},

	{
		"ReadIndividualBdtData",
		http.MethodGet,
		"/policy-data/bdt-data/:bdtReferenceId",
		ReadIndividualBdtData,
	},

	{
		"UpdateIndividualBdtData",
		http.MethodPatch,
		"/policy-data/bdt-data/:bdtReferenceId",
		UpdateIndividualBdtData,
	},

	{
		"DeleteIndividualPolicyDataSubscription",
		http.MethodDelete,
		"/policy-data/subs-to-notify/:subsId",
		DeleteIndividualPolicyDataSubscription,
	},

	{
		"ReplaceIndividualPolicyDataSubscription",
		http.MethodPut,
		"/policy-data/subs-to-notify/:subsId",
		ReplaceIndividualPolicyDataSubscription,
	},

	{
		"GetMBSSessPolCtrlData",
		http.MethodGet,
		"/policy-data/mbs-session-pol-data/:polSessionId",
		GetMBSSessPolCtrlData,
	},

	{
		"DeleteOperatorSpecificData",
		http.MethodDelete,
		"/policy-data/ues/:ueId/operator-specific-data",
		DeleteOperatorSpecificData,
	},

	{
		"ReadOperatorSpecificData",
		http.MethodGet,
		"/policy-data/ues/:ueId/operator-specific-data",
		ReadOperatorSpecificData,
	},

	{
		"ReplaceOperatorSpecificData",
		http.MethodPut,
		"/policy-data/ues/:ueId/operator-specific-data",
		ReplaceOperatorSpecificData,
	},

	{
		"UpdateOperatorSpecificData",
		http.MethodPatch,
		"/policy-data/ues/:ueId/operator-specific-data",
		UpdateOperatorSpecificData,
	},

	{
		"ReadPlmnUePolicySet",
		http.MethodGet,
		"/policy-data/plmns/:plmnId/ue-policy-set",
		ReadPlmnUePolicySet,
	},

	{
		"ReadPolicyData",
		http.MethodGet,
		"/policy-data/ues/:ueId",
		ReadPolicyData,
	},

	{
		"CreateIndividualPolicyDataSubscription",
		http.MethodPost,
		"/policy-data/subs-to-notify",
		CreateIndividualPolicyDataSubscription,
	},

	{
		"ReadSessionManagementPolicyData",
		http.MethodGet,
		"/policy-data/ues/:ueId/sm-data",
		ReadSessionManagementPolicyData,
	},

	{
		"UpdateSessionManagementPolicyData",
		http.MethodPatch,
		"/policy-data/ues/:ueId/sm-data",
		UpdateSessionManagementPolicyData,
	},

	{
		"ReadSlicePolicyControlData",
		http.MethodGet,
		"/policy-data/slice-control-data/:snssai",
		ReadSlicePolicyControlData,
	},

	{
		"UpdateSlicePolicyControlData",
		http.MethodPatch,
		"/policy-data/slice-control-data/:snssai",
		UpdateSlicePolicyControlData,
	},

	{
		"ReadSponsorConnectivityData",
		http.MethodGet,
		"/policy-data/sponsor-connectivity-data/:sponsorId",
		ReadSponsorConnectivityData,
	},

	{
		"CreateOrReplaceUEPolicySet",
		http.MethodPut,
		"/policy-data/ues/:ueId/ue-policy-set",
		CreateOrReplaceUEPolicySet,
	},

	{
		"ReadUEPolicySet",
		http.MethodGet,
		"/policy-data/ues/:ueId/ue-policy-set",
		ReadUEPolicySet,
	},

	{
		"UpdateUEPolicySet",
		http.MethodPatch,
		"/policy-data/ues/:ueId/ue-policy-set",
		UpdateUEPolicySet,
	},

	{
		"CreateUsageMonitoringResource",
		http.MethodPut,
		"/policy-data/ues/:ueId/sm-data/:usageMonId",
		CreateUsageMonitoringResource,
	},

	{
		"DeleteUsageMonitoringInformation",
		http.MethodDelete,
		"/policy-data/ues/:ueId/sm-data/:usageMonId",
		DeleteUsageMonitoringInformation,
	},

	{
		"ReadUsageMonitoringInformation",
		http.MethodGet,
		"/policy-data/ues/:ueId/sm-data/:usageMonId",
		ReadUsageMonitoringInformation,
	},
}
