package service

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	context2 "golang.org/x/net/context"
	"io/fs"
	"net/http"
	"testing"
	"time"
)

type Dummy struct {
	driver.Driver
}

func TestNewService(t *testing.T) {
	l := logrus.NewEntry(logrus.New())
	sql.Register("sample", Dummy{})
	db, err := sql.Open("sample", "sample")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	s := NewService(&Config{}, l, db, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}))

	assert.NotNil(t, s)
	assert.NotNil(t, s.API())
	assert.NotNil(t, s.Logger())
	assert.NotNil(t, s.Config())
	assert.Nil(t, s.RunMigration(func(db *sql.DB, schemaFs fs.FS, schema string) error {
		return nil
	}, nil, "sample-schema"))
}

func TestRun(t *testing.T) {
	l := logrus.NewEntry(logrus.New())
	s := NewService(&Config{Address: ":10"}, l, nil, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}))
	ctx, cancel := context2.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	type args struct {
		ctx context.Context
		s   Service
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "TestRun",
			args: args{
				ctx: ctx,
				s:   s,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, Run(tt.args.ctx, tt.args.s), fmt.Sprintf("Run(%v, %v)", tt.args.ctx, tt.args.s))
		})
	}
}
