//go:build tools
// +build tools

package main

import (
	_ "github.com/campoy/embedmd"
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/vektra/mockery/v2"
	_ "github.com/zitadel/oidc/v3/example/server"
)
