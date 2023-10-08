package logger

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var l *logrus.Entry

type ctKey struct{}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

// New returns a logrus o*log. Entry with the givenn fields and formatter
func New(fields logrus.Fields, level logrus.Level) *logrus.Entry {
	if l == nil {
		location, _ := time.LoadLocation("America/New_York")
		lg := logrus.StandardLogger()
		lg.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})
		lg.SetOutput(os.Stdout)
		lg.SetLevel(level)
		l = logrus.NewEntry(lg).WithFields(fields).WithFields(logrus.Fields{"timestamp": time.Now().UTC().In(location).Format(time.RFC3339Nano)})
		l.Level = level
		return l
	}

	l.Logger.SetLevel(level)
	l = logrus.NewEntry(l.Logger).WithFields(fields)
	l.Level = level
	return l
}

// From returns a logrus object (newLogger) with the added context
func From(ctx context.Context) *logrus.Entry {
	if l, ok := ctx.Value(ctKey{}).(*logrus.Entry); ok {
		return l
	}
	if l == nil {
		l = New(logrus.Fields{}, logrus.InfoLevel)
	}
	return l
}

// With returns a context with the added logrus.Entry
func With(ct context.Context, logger *logrus.Entry) context.Context {
	if l == nil {
		l = New(logrus.Fields{}, logrus.InfoLevel)
	}
	return context.WithValue(ct, ctKey{}, logger)
}

// WithFormatter returns a logrus object (newLogger) with the added formatter
func WithFormatter(formatter logrus.Formatter) *logrus.Entry {
	if l == nil {
		l = New(logrus.Fields{}, logrus.InfoLevel)
	}
	if formatter != nil {
		l.Logger.Formatter = formatter
	}
	return l
}

func Debug() bool {
	if l != nil {
		return l.Logger.Level == logrus.DebugLevel
	}
	return false
}

type stackTracer interface {
	StackStrace() errors.StackTrace
}

// WithError - prints the stack-strace of the error
func WithError(ctx context.Context, err error) *logrus.Entry {
	n := l
	if n == nil {
		n = New(logrus.Fields{}, logrus.InfoLevel)
	}
	if stackErr, ok := err.(stackTracer); ok {
		n = From(ctx).WithField("stacktrace", stackErr.StackStrace())
	}
	if err != nil {
		return n.WithError(err)
	}
	return n
}
