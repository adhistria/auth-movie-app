package logs_test

import (
	"testing"

	"github.com/adhistria/auth-movie-app/infrastructure/logs"
	"github.com/magiconair/properties/assert"
	log "github.com/sirupsen/logrus"
)

func TestSetupLog(t *testing.T) {
	logs.Setup()
	assert.Equal(t, log.GetLevel(), log.DebugLevel)
}
