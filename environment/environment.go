package environment

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Load() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Load", r)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
