package api

import (
	"cerberus/app/institution"
	"cerberus/services/logger"

	"github.com/gin-gonic/gin"
)

func GetInstitutionAccountData(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	data, err := institution.GetAccountById(id)

	c.Header("Content-Type", "application/json")

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountHistory(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	data, err := institution.GetAccountHistory(id)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountDocumentData(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, err := institution.GetAccountDocument(id, documentName)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountDocumentVersion(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	version := c.Request.URL.Query().Get("version")

	data, err := institution.GetAccountDocumentVersion(id, documentName, version)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountDocumentVersions(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, err := institution.GetAccountDocumentVersions(id, documentName)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountsByEmail(c *gin.Context) {

	email := c.Request.URL.Query().Get("email")

	data, err := institution.GetAccountsByEmail(email)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountsByOrgName(c *gin.Context) {

	name := c.Request.URL.Query().Get("name")

	data, err := institution.GetAccountsByOrgName(name)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountsByContactPerson(c *gin.Context) {

	contactPerson := c.Request.URL.Query().Get("contactperson")

	data, err := institution.GetAccountsByContactPerson(contactPerson)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

func GetInstitutionAccountsBySelector(c *gin.Context) {

	selectorKey := c.Request.URL.Query().Get("selector")
	value := c.Request.URL.Query().Get("value")

	data, err := institution.GetAccountsBySelector(selectorKey, value)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}
