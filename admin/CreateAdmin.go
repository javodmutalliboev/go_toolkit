package admin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/javodmutalliboev/go_toolkit/interface_package"
	"github.com/javodmutalliboev/go_toolkit/password"
	"github.com/javodmutalliboev/go_toolkit/struct_package"
)

func CreateAdmin() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to create an admin? (yes/no)")
	response, _ := reader.ReadString('\n')

	response = strings.ToLower(strings.TrimSpace(response))

	if response == "yes" {
		// Code to create an admin goes here
		fmt.Println("Creating an admin...")

		var admin interface_package.Admin = &struct_package.Admin{}

		fmt.Println("Enter name:")
		name, _ := reader.ReadString('\n')
		admin.(*struct_package.Admin).Name = strings.TrimSpace(name)

		fmt.Println("Enter surname:")
		surname, _ := reader.ReadString('\n')
		admin.(*struct_package.Admin).Surname = strings.TrimSpace(surname)

		fmt.Println("Enter email:")
		email, _ := reader.ReadString('\n')
		admin.(*struct_package.Admin).Email = strings.TrimSpace(email)

		fmt.Println("Enter password:")
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimSpace(pass)
		hashedPassword, err := password.HashPassword(pass)
		if err != nil {
			fmt.Println("Error hashing password.")
			return
		}
		admin.(*struct_package.Admin).Password = hashedPassword

		err = admin.CreateAdmin()
		if err != nil {
			fmt.Printf("Error creating admin. %s", err)
			return
		}

		fmt.Println("Admin created successfully.")
	} else if response == "no" {
		fmt.Println("Not creating an admin.")
	} else {
		fmt.Println("Invalid response. Please answer with 'yes' or 'no'.")
	}
}
