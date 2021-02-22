package config

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func setSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://72dc3767dd9a4c3184f8880556dd58fd@o528515.ingest.sentry.io/5645909",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works caca!")
}
