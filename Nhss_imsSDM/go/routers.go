/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nhss_imsSDM

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
		"/nhss-ims-sdm/v1/",
		Index,
	},

	{
		"GetLocCsDomain",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/cs-domain/location-data",
		GetLocCsDomain,
	},

	{
		"GetCsrn",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/cs-domain/csrn",
		GetCsrn,
	},

	{
		"GetCsUserStateInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/cs-domain/user-state",
		GetCsUserStateInfo,
	},

	{
		"GetChargingInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/profile-data/charging-info",
		GetChargingInfo,
	},

	{
		"GetDsaiInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/dsai",
		GetDsaiInfo,
	},

	{
		"DeleteRepositoryDataServInd",
		http.MethodDelete,
		"/nhss-ims-sdm/v1/:imsUeId/repository-data/:serviceIndication",
		DeleteRepositoryDataServInd,
	},

	{
		"DeleteSmsRegistrationInfo",
		http.MethodDelete,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/sms-registration-info",
		DeleteSmsRegistrationInfo,
	},

	{
		"GetIfcs",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/profile-data/ifcs",
		GetIfcs,
	},

	{
		"GetIMEISVInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/identities/imeisv",
		GetIMEISVInfo,
	},

	{
		"GetProfileData",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/profile-data",
		GetProfileData,
	},

	{
		"GetIpAddressInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/ip-address",
		GetIpAddressInfo,
	},

	{
		"GetLocPsDomain",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/location-data",
		GetLocPsDomain,
	},

	{
		"GetPsUserStateInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/user-state",
		GetPsUserStateInfo,
	},

	{
		"GetPriorityInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/profile-data/priority-levels",
		GetPriorityInfo,
	},

	{
		"UeReachUnsubscribe",
		http.MethodDelete,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/ue-reach-subscriptions/:subscriptionId",
		UeReachUnsubscribe,
	},

	{
		"UeReachSubsModify",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/ue-reach-subscriptions/:subscriptionId",
		UeReachSubsModify,
	},

	{
		"GetReferenceLocationInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/wireline-domain/reference-location",
		GetReferenceLocationInfo,
	},

	{
		"GetRegistrationStatus",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/registration-status",
		GetRegistrationStatus,
	},

	{
		"GetRepositoryDataServInd",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/repository-data/:serviceIndication",
		GetRepositoryDataServInd,
	},

	{
		"GetRepositoryDataServIndList",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/repository-data",
		GetRepositoryDataServIndList,
	},

	{
		"GetImsPrivateIds",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/identities/private-identities",
		GetImsPrivateIds,
	},

	{
		"GetImsAssocIds",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/identities/ims-associated-identities",
		GetImsAssocIds,
	},

	{
		"GetPsiState",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/psi-status",
		GetPsiState,
	},

	{
		"GetSharedData",
		http.MethodGet,
		"/nhss-ims-sdm/v1/shared-data",
		GetSharedData,
	},

	{
		"GetMsisdns",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/identities/msisdns",
		GetMsisdns,
	},

	{
		"GetScscfCapabilities",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/location-data/scscf-capabilities",
		GetScscfCapabilities,
	},

	{
		"GetScscfSelectionAssistanceInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/location-data/scscf-selection-assistance-info",
		GetScscfSelectionAssistanceInfo,
	},

	{
		"GetSrvccData",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/srvcc-data",
		GetSrvccData,
	},

	{
		"ImsSdmSubscribe",
		http.MethodPost,
		"/nhss-ims-sdm/v1/:imsUeId/subscriptions",
		ImsSdmSubscribe,
	},

	{
		"ImsSdmUnsubscribe",
		http.MethodDelete,
		"/nhss-ims-sdm/v1/:imsUeId/subscriptions/:subscriptionId",
		ImsSdmUnsubscribe,
	},

	{
		"ImsSdmSubsModify",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/:imsUeId/subscriptions/:subscriptionId",
		ImsSdmSubsModify,
	},

	{
		"GetSmsRegistrationInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/sms-registration-info",
		GetSmsRegistrationInfo,
	},

	{
		"GetServerName",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/location-data/server-name",
		GetServerName,
	},

	{
		"GetServiceTraceInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/ims-data/profile-data/service-level-trace-information",
		GetServiceTraceInfo,
	},

	{
		"SubscribeToSharedData",
		http.MethodPost,
		"/nhss-ims-sdm/v1/shared-data-subscriptions",
		SubscribeToSharedData,
	},

	{
		"UnsubscribeForSharedData",
		http.MethodDelete,
		"/nhss-ims-sdm/v1/shared-data-subscriptions/:subscriptionId",
		UnsubscribeForSharedData,
	},

	{
		"ModifySharedDataSubs",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/shared-data-subscriptions/:subscriptionId",
		ModifySharedDataSubs,
	},

	{
		"GetTadsInfo",
		http.MethodGet,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/tads-info",
		GetTadsInfo,
	},

	{
		"UeReachIpSubscribe",
		http.MethodPost,
		"/nhss-ims-sdm/v1/:imsUeId/access-data/ps-domain/ue-reach-subscriptions",
		UeReachIpSubscribe,
	},

	{
		"UpdateDsaiState",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/dsai",
		UpdateDsaiState,
	},

	{
		"UpdatePsiState",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/psi-status",
		UpdatePsiState,
	},

	{
		"UpdateRepositoryDataServInd",
		http.MethodPut,
		"/nhss-ims-sdm/v1/:imsUeId/repository-data/:serviceIndication",
		UpdateRepositoryDataServInd,
	},

	{
		"UpdateSmsRegistrationInfo",
		http.MethodPut,
		"/nhss-ims-sdm/v1/:imsUeId/service-data/sms-registration-info",
		UpdateSmsRegistrationInfo,
	},

	{
		"UpdateSrvccData",
		http.MethodPatch,
		"/nhss-ims-sdm/v1/:imsUeId/srvcc-data",
		UpdateSrvccData,
	},
}