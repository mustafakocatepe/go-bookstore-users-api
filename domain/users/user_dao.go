package users

import (
	"fmt"
	"github.com/mustafakocatepe/go-bookstore-users-api/datasources/mysql/users_db"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/date_utils"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//stmt.QueryRow() //Eğer sorgumuz bize tek bir row döndürecekse QueryRow kullanırız. Geri dönüş tipi *Row
	//stmt.Query()    //Eğer sorgumuz bize birden fazla row döndürecekse Query kullanırız. Geri dönüş tipi *Rows

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.NewBadRequestError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
