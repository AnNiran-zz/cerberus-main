package api

import (
	"cerberus/app/person"
	"cerberus/services/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCreatePersonAccount(c *gin.Context) {

	firstName := c.Request.URL.Query().Get("firstname")
	lastName := c.Request.URL.Query().Get("lastname")
	email := c.Request.URL.Query().Get("email")
	phone := c.Request.URL.Query().Get("phone")

	_, data, err := person.CreateAccount(firstName, lastName, email, phone)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountBySelector(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	selector := c.Request.URL.Query().Get("selector")
	value := c.Request.URL.Query().Get("value")

	data, _, err := person.UpdateAccountBySelector(id, selector, value)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountFirstName(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	firstName := c.Request.URL.Query().Get("firstname")

	data, _, err := person.UpdateAccountFirstName(id, firstName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountLastName(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	lastName := c.Request.URL.Query().Get("lastname")

	data, _, err := person.UpdateAccountFirstName(id, lastName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountPhone(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	phone := c.Request.URL.Query().Get("phne")

	data, _, err := person.UpdateAccountPhone(id, phone)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountEmail(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	email := c.Request.URL.Query().Get("email")

	data, _, err := person.UpdateAccountEmail(id, email)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostDeletePersonAccount(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	_, err := person.DeleteAccount(id)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, nil)
}

func PostCreatePersonAccountDocument(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	holderName := c.Request.URL.Query().Get("holder")
	countryIssue := c.Request.URL.Query().Get("countryissue")
	filename := c.Request.URL.Query().Get("filename")

	_, _, rsaLink, err := person.CreateNewDocument(id, documentName, holderName, countryIssue, filename)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	// TODO: figure out how to return rsa link and the data in the response
	c.JSON(200, rsaLink)
}

func PostCreatePersonAccountDocumentVersion(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	filename := c.Request.URL.Query().Get("filename")
	rsaLink := c.Request.URL.Query().Get("rsalink")

	data, _, err := person.CreateDocumentVersion(id, documentName, filename, rsaLink)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostUpdatePersonAccountDocumentCountryIssue(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	countryIssueUpdate := c.Request.URL.Query().Get("country")

	data, _, err := person.UpdateDocumentCountryIssue(id, documentName, countryIssueUpdate)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)

}

func PostUpdatePersonAccountDocumentHolderName(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	holder := c.Request.URL.Query().Get("holder")

	data, _, err := person.UpdateDocumentHolderName(id, documentName, holder)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostDeletePersonAccountDocument(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, _, err := person.DeleteDocument(id, documentName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func PostDeletePersonAccountDocumentVersion(c *gin.Context) {

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

	data, _, err := person.DeleteDocumentVersion(id, documentName, versionAsString)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)

}
