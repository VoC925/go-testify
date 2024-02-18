package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigNew(t *testing.T) {
	pathToConfig := `config.yml`

	cnf := New(pathToConfig)
	require.NotNil(t, *cnf)
}
