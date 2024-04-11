package struct_package

import (
	"fmt"

	"github.com/javodmutalliboev/go_toolkit/password"
	"github.com/javodmutalliboev/go_toolkit/postgresql"
)

type Admin struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (a *Admin) Login() error {
	db, err := postgresql.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	pass := a.Password

	err = db.QueryRow("SELECT id, email, password, name, surname, created_at, updated_at FROM admin WHERE email = $1", a.Email).Scan(&a.ID, &a.Email, &a.Password, &a.Name, &a.Surname, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return err
	}

	authenticated := password.CheckPasswordHash(pass, a.Password)
	if !authenticated {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func (admin *Admin) Logout() error {
	return nil
}

func (a *Admin) CreateAdmin() error {
	db, err := postgresql.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO admin (email, password, name, surname) VALUES ($1, $2, $3, $4)", a.Email, a.Password, a.Name, a.Surname)
	if err != nil {
		return err
	}

	return nil
}
