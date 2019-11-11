package main

import (
	institution "cerberus/api/accounts/institution"
	person "cerberus/api/accounts/person"
	"cerberus/services/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Log("api_endpoints")
	logger.LogEvent("Listening for api requests", nil)

	// gin
	router := gin.New()

	personAccnts := router.Group("/person")
	institutionAccnts := router.Group("/institution")

	// person: query single account data
	// http://159.203.35.156:8000/person/get/greeting?name=vito
	personAccnts.GET("/get/greeting", person.GetGreeting)
	personAccnts.GET("/get/account/data", person.GetPersonAccountData)
	personAccnts.GET("/get/account/history", person.GetPersonAccountHistory)
	personAccnts.GET("/get/document/data", person.GetPersonAccountDocumentData)
	personAccnts.GET("/get/document/version", person.GetPersonAccountDocumentVersion)
	personAccnts.GET("/get/document/versions", person.GetPersonAccountDocumentVersions)

	// person: query multiple accounts data
	personAccnts.GET("/get/accounts/email", person.GetPersonAccountsByEmail)
	personAccnts.GET("/get/accounts/firstname", person.GetPersonAccountsByFirstName)
	personAccnts.GET("/get/accounts/lastname", person.GetPersonAccountsByLastName)

	// person: invoke
	// http://159.203.35.156:8000/person/create/account
	//
	personAccnts.POST("/create/account", person.PostCreatePersonAccount) // ok
	personAccnts.POST("/update/account/selector", person.PostUpdatePersonAccountBySelector)
	personAccnts.POST("/update/account/firstname", person.PostUpdatePersonAccountFirstName)
	personAccnts.POST("/update/account/lastname", person.PostUpdatePersonAccountLastName)
	personAccnts.POST("/update/account/phone", person.PostUpdatePersonAccountPhone)
	personAccnts.POST("/update/account/email", person.PostUpdatePersonAccountEmail)
	personAccnts.POST("/delete/account", person.PostDeletePersonAccount)

	personAccnts.POST("/create/document/new", person.PostCreatePersonAccountDocument)
	personAccnts.POST("/create/document/version", person.PostCreatePersonAccountDocumentVersion)
	personAccnts.POST("/update/document/countryissue", person.PostUpdatePersonAccountDocumentCountryIssue)
	personAccnts.POST("/update/document/holder", person.PostUpdatePersonAccountDocumentHolderName)
	personAccnts.POST("/delete/document", person.PostDeletePersonAccountDocument)
	personAccnts.POST("/delete/document/version", person.PostDeletePersonAccountDocumentVersion)

	// institution: query single account data
	institutionAccnts.GET("/get/account/data", institution.GetInstitutionAccountData)
	institutionAccnts.GET("/get/account/history", institution.GetInstitutionAccountHistory)
	institutionAccnts.GET("/get/document/data", institution.GetInstitutionAccountDocumentData)
	institutionAccnts.GET("/get/document/version", institution.GetInstitutionAccountDocumentVersion)
	institutionAccnts.GET("/get/document/versions", institution.GetInstitutionAccountDocumentVersions)

	// institution: query multiple accounts data
	institutionAccnts.GET("/get/accounts/email", institution.GetInstitutionAccountsByEmail)
	institutionAccnts.GET("/get/accounts/name", institution.GetInstitutionAccountsByOrgName)
	institutionAccnts.GET("/get/accounts/contact", institution.GetInstitutionAccountsByContactPerson)
	institutionAccnts.GET("/get/accounts/selector", institution.GetInstitutionAccountsBySelector)

	// institution: invoke
	institutionAccnts.POST("/create/account", institution.PostCreateInstitutionAccount)
	institutionAccnts.POST("/update/account/selector", institution.PostUpdateInstitutionAccountBySelector)
	institutionAccnts.POST("/update/account/name", institution.PostUpdateInstitutionAccountName)
	institutionAccnts.POST("/update/account/email", institution.PostUpdateInstitutionAccountEmail)
	institutionAccnts.POST("/update/account/phone", institution.PostUpdateInstitutionAccountPhone)
	institutionAccnts.POST("/update/account/contact", institution.PostUpdateInstitutionAccountContactPerson)
	institutionAccnts.POST("/update/account/address", institution.PostUpdateInstitutionAccountAddress)
	institutionAccnts.POST("/delete/account", institution.PostDeleteInstitutionAccount)

	institutionAccnts.POST("/create/document/new", institution.PostCreateInstitutionAccountDocument)
	institutionAccnts.POST("/create/document/version", institution.PostCreateInstitutionAccountDocumentVersion)
	institutionAccnts.POST("/update/document/countryissue", institution.PostUpdateInstitutionAccountDocumentCountryIssue)
	institutionAccnts.POST("/update/document/holder", institution.PostUpdateInstitutionAccountDocumentHolderName)
	institutionAccnts.POST("/delete/document", institution.PostDeleteInstitutionAccountDocument)
	institutionAccnts.POST("/delete/document/version", institution.PostDeleteInstitutionAccountDocumentVersion)

	err := router.Run(":8000")

	if err != nil {
		logger.LogEvent("Server stopped because of error", err)
	}
}
