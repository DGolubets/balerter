package alertmanager

import (
	"github.com/balerter/balerter/internal/config/channels/alertmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	a, err := New(alertmanager.Alertmanager{}, nil)
	require.NoError(t, err)
	assert.IsType(t, &AlertManager{}, a)
}

func TestName(t *testing.T) {
	a := &AlertManager{name: "foo"}
	assert.Equal(t, "foo", a.Name())
}