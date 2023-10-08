package logger

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	log := New(logrus.Fields{}, logrus.DebugLevel)
	assert.NotNil(t, log)
	assert.Equal(t, logrus.DebugLevel, log.Level)
	assert.Equal(t, logrus.DebugLevel, log.Logger.Level)
}

func TestFrom(t *testing.T) {
	log := From(context.Background())
	assert.NotNil(t, log)
}

func TestWith(t *testing.T) {
	log := New(logrus.Fields{}, logrus.DebugLevel)
	ctx := With(context.Background(), log)
	assert.NotNil(t, ctx)
	log2 := From(context.Background())
	assert.Equal(t, log, log2)
}

func TestWithFormatter(t *testing.T) {
	f := &logrus.JSONFormatter{}
	log := WithFormatter(f)
	assert.NotNil(t, log)
	assert.Equal(t, f, log.Logger.Formatter)
}

func TestWithError(t *testing.T) {
	expErr := errors.New("test error")
	log := WithError(context.Background(), expErr)

	assert.NotNil(t, log)
	s, err := log.String()
	assert.Nil(t, err)
	assert.True(t, strings.Contains(s, expErr.Error()))
}

func TestDebug(t *testing.T) {
	_ = New(logrus.Fields{}, logrus.InfoLevel)
	assert.False(t, Debug())
	_ = New(logrus.Fields{}, logrus.DebugLevel)
	assert.True(t, Debug())
}
