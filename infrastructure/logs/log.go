package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Setup set format, output, and level of logs
func Setup() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
