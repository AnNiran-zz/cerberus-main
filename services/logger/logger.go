package logger

import (
	"cerberus/config"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type eventLog struct {
	id           bson.ObjectId
	message      string
	errorMessage string
	error        error
}

func Log(filename string) {

	// TODO: check if provided filename contains spaces

	// logfile -> create
	logfile, err := os.OpenFile(config.LogsPath+filename+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)

	if err != nil {
		fmt.Printf("Failed to create %s.log file", filename)
		os.Exit(1)
	}

	//defer logfile.Close()

	currentDateTime := getTime()

	log.Println(currentDateTime + ": server log started")
	log.SetOutput(logfile)
}

func LogEvent(message string, err error) {

	var event = &eventLog{
		id: bson.NewObjectId(),
	}

	if message != "" {
		event.message = " " + message
		log.Println(event.id.Hex() + event.message)
	}

	if err != nil {
		event.errorMessage = " Error: " + err.Error()
		event.error = err

		log.Println(event.id.Hex() + event.errorMessage)
	}
}

func getTime() string {

	currentDateTime := time.Now()
	CurrentDateTime := currentDateTime.Format("2006-01-02 15:04:05")

	return CurrentDateTime
}
