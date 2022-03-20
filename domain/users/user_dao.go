package users

import (
	"fmt"
	"github.com/mustafakocatepe/go-bookstore-users-api/datasources/mysql/users_db"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/errors"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare("DELETE FROM users WHERE id=?;")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil { //Donen result ile ilgilenmiyorum. Error'den zaten bir hata olup olmadigin anlayabiliyorum.
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare("SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;")
	if err != nil {
		//logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status) // ? = status
	if err != nil {
		//logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close() // err olup olmadigini kontrol ettikten sonra rows.close yapilmali. Kontrol isleminden once yapilirsa rows'un nil olma ihtimali vardir.

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			//logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
