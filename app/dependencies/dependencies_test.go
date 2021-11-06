package dependencies_test

import (
	"os"
	"testing"

	"github.com/blitzshare/blitzshare.bootstrap.node/app/config"
	"github.com/blitzshare/blitzshare.bootstrap.node/app/dependencies"
	"github.com/stretchr/testify/assert"
)

func TestDependencies(t *testing.T) {
	setUp()

	cfg, err := config.Load()
	assert.Nil(t, err, "config should load")
	deps, err := dependencies.NewDependencies(cfg)
	assert.Nil(t, err, "dependencies should be initialised")
	assert.NotNilf(t, deps.Config, "config should be loaded in dependencies")

	tearDown()
}

func setUp() {
	_ = os.Setenv("PORT", "38643")
	_ = os.Setenv("ENV", "local")
}

func tearDown() {
	_ = os.Unsetenv("PORT")
	_ = os.Unsetenv("ENV")
}
