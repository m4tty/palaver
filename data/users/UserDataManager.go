package data

type UserDataManager interface {
	GetUsers() (results *[]User, err error)
	GetUserById(id string) (result User, err error)
	SaveUser(user *User) (key string, err error)
	DeleteUser(id string) (err error)
}
