package services

import (
	"github.com/mustafakocatepe/go-bookstore-users-api/domain/users"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/date_utils"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) { //Geri dönüş tipi olarak users.User olarak dönersek nil dönemiyoruz. Fakat geri dönüş tipi *users.User olursa nil dönülebilir.
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUSer(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
	/* users, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return users, nil*/

}
