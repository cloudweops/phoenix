package rest_test

import (
	"context"
	"testing"

	"github.com/cloudweops/phoenix/logger/zap"
	"{{.PKG}}/clients/rest"
	"{{.PKG}}/test/tools"
)

var (
	client *rest.ClientSet
	ctx    = context.Background()
)

func init() {
	zap.DevelopmentSetup()
	conf := rest.NewDefaultConfig()
	client = rest.NewClient(conf)
}