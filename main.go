package main

import (
	"github.com/javodmutalliboev/go_toolkit/environment"
	"github.com/javodmutalliboev/go_toolkit/xlsx"
)

func init() {
	environment.Load()
}

func main() {
	// admin.CreateAdmin()
	// session.GetSession2()
	xlsx.Export()
}
