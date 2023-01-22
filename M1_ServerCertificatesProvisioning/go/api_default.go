/*
 * M1_ServerCertificatesProvisioning
 *
 * 5GMS AF M1 Server Certificates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package M1_ServerCertificatesProvisioning

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrReserveServerCertificate - Create or reserve a Service Certificate resource
func CreateOrReserveServerCertificate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DestroyServerCertificate - Destroy an existing Server Certificate resource
func DestroyServerCertificate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// RetrieveServerCertificate - Retrieve the X.509 certificate representation of the specified Server Certificate resource
func RetrieveServerCertificate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UploadServerCertificate - Upload the X.509 certificate for a previously reserved Server Certificate resource
func UploadServerCertificate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
