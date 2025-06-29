package clients

import (
	"fmt"

	"github.com/maxabella/appgym/dao"
)

type User interface {
	GetUserById()
	CreateUser(dao.User) (uint, error)
	GetUserByEmail(string) (dao.User, error)
}

type UserClient struct {
	DbClient *Mysql_Client
}

func (u *UserClient) GetUserById(id int) (dao.User, error) {
	var user dao.User
	if err := u.DbClient.db.First(&user, "id = ?", id); err != nil {
		return dao.User{}, err.Error
	}
	return user, nil
}

func (u *UserClient) GetUserByEmail(email string) (dao.User, error) {
	var user dao.User
	if err := u.DbClient.db.First(&user, "email = ?", email); err.Error != nil {
		fmt.Println("second attempt ahh: ", user.HashedPassword)
		return dao.User{}, err.Error
	}
	fmt.Println("second atempt ahh2: ", user.HashedPassword)
	return user, nil
}

func (u *UserClient) CreateUser(user dao.User) (uint, uint, error) {
	result := u.DbClient.db.Create(&user)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return user.ID, user.UserTypeID, nil // returns the user without the token and no error
}
