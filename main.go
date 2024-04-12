package main

import (
	"github.com/javodmutalliboev/go_toolkit/environment"
	"github.com/javodmutalliboev/go_toolkit/session"
)

func init() {
	environment.Load()
}

func main() {
	// admin.CreateAdmin()
	session.GetSession2()
}
