package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	databaseAddressTest = "mysql://localhost:3306/dbname"

	serverHostTest = "localhost"
	serverPortTest = "8080"

	serverAddressTest = "localhost:8080"
)

var expectedCfg = Config{
	Database: Database{Address: databaseAddressTest},
	Server: Server{
		Host: serverHostTest,
		Port: serverPortTest,
	},
}

func TestLoad(t *testing.T) {
	err := os.Setenv(databaseAddressEnv, databaseAddressTest)
	require.NoError(t, err)

	err = os.Setenv(serverHostEnv, serverHostTest)
	require.NoError(t, err)

	err = os.Setenv(serverPortEnv, serverPortTest)
	require.NoError(t, err)

	cfg := Load()

	assert.Equal(t, expectedCfg, cfg)

	assert.Equal(t, serverAddressTest, cfg.GetServerAddress())
}
