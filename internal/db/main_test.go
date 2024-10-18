package db

import (
	"os"
	"testing"

	"github.com/longln/reboot-simplebank/global"
	"github.com/longln/reboot-simplebank/internal/initialize"
)

var (
	testQueries *Queries
)

func TestMain(m *testing.M) {
	initialize.Run()
	testQueries = New(global.Database)
	os.Exit(m.Run())
}