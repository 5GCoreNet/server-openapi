/*
 * 3gpp-chargeable-party
 *
 * API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ChargeableParty

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateChargeablePartyTransaction - Create a new chargeable party transaction resource.
func CreateChargeablePartyTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchAllChargeablePartyTransactions - Read all or queried chargeable party transaction resources for a given SCS/AS.
func FetchAllChargeablePartyTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
