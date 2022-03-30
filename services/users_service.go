package services

import (
	"github.com/mustafakocatepe/go-bookstore-users-api/domain/users"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/crypto_utils"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/date_utils"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{} //Interface icerisinde func.'larimizi bu package icerisinde tanimlamais olsaydik. implementasyon hatasini bu satirda alirdik.
)

type usersService struct{} // userService adinda bir struct olusturduk.
// Daha sonrasinda yukarida bu strcut'tan bir variable tanimladik.
// func. basina bu struct'i vererek receiver fun. olmasini sagladik.
//controller da ise cagirirken tanimladigimiz variable sayesinde services.UsersService diyerek de func.'larimiza erisebildik.

type usersServiceInterface interface { //interface kullanimlarinda sadece data type yazmamiz yeterli olabiliyor. isimlendirme yapmasak da olur.
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	//LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) { //Geri dönüş tipi olarak users.User olarak dönersek nil dönemiyoruz. Fakat geri dönüş tipi *users.User olursa nil dönülebilir.
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := UsersService.GetUser(user.Id)
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

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
	/* users, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return users, nil*/

}
