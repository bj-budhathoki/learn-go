package main

import (
	"github.com/bj-budhathoki/learn-go/url_short/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
