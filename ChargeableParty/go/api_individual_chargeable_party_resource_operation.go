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

// DeleteChargeablePartyTransaction - Deletes a chargeable party resource for a given SCS/AS and a transcation Id.
func DeleteChargeablePartyTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// FetchIndChargeablePartyTransaction - Read a chargeable party resource for a given SCS/AS and a transaction Id.
func FetchIndChargeablePartyTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateChargeablePartyTransaction - Updates a existing chargeable party resource for a given SCS/AS and transaction Id.
func UpdateChargeablePartyTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
