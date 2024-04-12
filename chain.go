package main

import (
	"net/http"

	"github.com/javodmutalliboev/go_toolkit/type_package"
)

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...type_package.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
