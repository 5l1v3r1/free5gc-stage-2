/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/src/udr/udr_handler/udr_message"
)

// GetIdentityData - Retrieve identity data by SUPI or GPSI
func GetIdentityData(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventGetIdentityData, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}