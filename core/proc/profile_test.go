package proc

import (
	"strings"
	"testing"

	"github.com/mikokutou1/go-zero-m/core/logx/logtest"
	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	c := logtest.NewCollector(t)
	profiler := StartProfile()
	// start again should not work
	assert.NotNil(t, StartProfile())
	profiler.Stop()
	// stop twice
	profiler.Stop()
	assert.True(t, strings.Contains(c.String(), ".pprof"))
}
