package repository

type AuthRepository interface {
	Register()
	Login()
	Refresh()
}
