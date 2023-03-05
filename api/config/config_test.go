package config_test

import (
	"github.com/AltheaIX/PembayaranSPP/config"
	"testing"
)

func TestLoadDatabaseEnv(t *testing.T) {
	model := &config.Config{}
	model.LoadDatabaseEnv()
	t.Log(model.Database.Username)
}
