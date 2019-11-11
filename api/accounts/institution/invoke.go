package api

import (
	"cerberus/app/institution"
	"cerberus/services/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCreateInstitutionAccount(c *gin.Context) {

	organizationName := c.Request.URL.Query().Get("name")
	contactPerson := c.Request.URL.Query().Get("contact")
	address := c.Request.URL.Query().Get("address")
	email := c.Request.URL.Query().Get("email")
	phone := c.Request.URL.Query().Get("phone")

	_, data, err := institution.CreateAccount(organizationName, contactPerson, address, email, phone)

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

func PostUpdateInstitutionAccountBySelector(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	selector := c.Request.URL.Query().Get("selector")
	value := c.Request.URL.Query().Get("value")

	data, _, err := institution.UpdateAccountBySelector(id, selector, value)

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

func PostUpdateInstitutionAccountName(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	name := c.Request.URL.Query().Get("name")

	data, _, err := institution.UpdateAccountName(id, name)

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

func PostUpdateInstitutionAccountEmail(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	email := c.Request.URL.Query().Get("email")

	data, _, err := institution.UpdateAccountEmail(id, email)

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

func PostUpdateInstitutionAccountPhone(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	phone := c.Request.URL.Query().Get("phone")

	data, _, err := institution.UpdateAccountPhone(id, phone)

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

func PostUpdateInstitutionAccountContactPerson(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	contactPerson := c.Request.URL.Query().Get("contact")

	data, _, err := institution.UpdateAccountContactPerson(id, contactPerson)

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

func PostUpdateInstitutionAccountAddress(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	address := c.Request.URL.Query().Get("address")

	data, _, err := institution.UpdateAccountAddress(id, address)

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

func PostDeleteInstitutionAccount(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	_, err := institution.DeleteAccount(id)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, nil)
}

func PostCreateInstitutionAccountDocument(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	holder := c.Request.URL.Query().Get("holder")
	countryIssue := c.Request.URL.Query().Get("countryissue")
	filename := c.Request.URL.Query().Get("filename")

	_, rsaLink, _, err := institution.CreateNewDocument(id, documentName, holder, countryIssue, filename)

	if err != nil {
		// log error and return it
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, rsaLink)
}

func PostCreateInstitutionAccountDocumentVersion(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	filename := c.Request.URL.Query().Get("filename")
	rsaLink := c.Request.URL.Query().Get("rsalink")

	data, _, err := institution.CreateDocumentVersion(id, documentName, filename, rsaLink)

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

func PostUpdateInstitutionAccountDocumentCountryIssue(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	country := c.Request.URL.Query().Get("country")

	data, _, err := institution.UpdateDocumentCountryIssue(id, documentName, country)

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

func PostUpdateInstitutionAccountDocumentHolderName(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	holder := c.Request.URL.Query().Get("holder")

	data, _, err := institution.UpdateDocumentHolderName(id, documentName, holder)

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

func PostDeleteInstitutionAccountDocument(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, _, err := institution.DeleteDocument(id, documentName)

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

func PostDeleteInstitutionAccountDocumentVersion(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	documentVersion := c.Request.URL.Query().Get("version")

	versionAsString, err := strconv.Atoi(documentVersion)

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(500, err.Error())
		c.Abort()
		return
	}

	data, _, err := institution.DeleteDocumentVersion(id, documentName, versionAsString)

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
