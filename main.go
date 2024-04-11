package main

import (
	"github.com/javodmutalliboev/go_toolkit/admin"
	"github.com/javodmutalliboev/go_toolkit/environment"
)

func init() {
	environment.Load()
}

func main() {
	admin.CreateAdmin()
}
