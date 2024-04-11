package interface_package

type Admin interface {
	Login() error
	Logout() error
	CreateAdmin() error
}
