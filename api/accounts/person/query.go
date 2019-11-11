package api

import (
	"cerberus/app/person"
	"cerberus/services/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetGreeting(c *gin.Context) {

	name := c.Request.URL.Query().Get("name")

	fmt.Println("Hello from the server " + name)

	c.Header("Content-Type", "application/json")
	c.JSON(200, "hello, this is a response")
}

func GetPersonAccountData(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	data, err := person.GetAccountById(id)

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

func GetPersonAccountHistory(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")

	data, err := person.GetAccountHistory(id)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountDocumentData(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, err := person.GetAccountDocument(id, documentName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountDocumentVersion(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")
	version := c.Request.URL.Query().Get("version")

	data, err := person.GetAccountDocumentVersion(id, documentName, version)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountDocumentVersions(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	documentName := c.Request.URL.Query().Get("documentname")

	data, err := person.GetAccountDocumentVersions(id, documentName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountsByEmail(c *gin.Context) {

	email := c.Request.URL.Query().Get("email")

	data, err := person.GetAccountsByEmail(email)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountsByFirstName(c *gin.Context) {

	firstName := c.Request.URL.Query().Get("firstname")

	data, err := person.GetAccountsByFirstName(firstName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}

func GetPersonAccountsByLastName(c *gin.Context) {

	lastName := c.Request.URL.Query().Get("lastname")

	data, err := person.GetAccountsByFirstName(lastName)

	c.Header("Content-Type", "application/json")

	if err != nil {
		logger.LogEvent(err.Error(), err)
		c.JSON(200, err.Error())
		c.Abort()
		return
	}

	c.JSON(200, data)
}
