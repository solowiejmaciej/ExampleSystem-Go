package initializers

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func ConfigureLogs() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if os.Getenv("ENV") == "production" {
		f, err := os.Create("logs.log")
		if err != nil {
			log.Error("Error creating log file")
			log.Fatal(err)
		}
		multiWriter := io.MultiWriter(os.Stdout, f)
		log.SetOutput(multiWriter)
	}

	log.SetLevel(log.TraceLevel)
	log.Infof("Application started...")
}
